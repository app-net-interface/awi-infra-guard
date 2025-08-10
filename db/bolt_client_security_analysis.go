// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/app-net-interface/awi-infra-guard/types"
)

// SyncVpcPostProcess performs security analysis on all VPC resources
// This function runs after SyncVPCIndexes to populate security status fields
func (client *boltClient) SyncVpcPostProcess() error {
	fmt.Println("Starting VPC security post-processing...")

	// Get all VPC indexes to process
	vpcIndexes, err := client.ListVPCIndex()
	if err != nil {
		return fmt.Errorf("failed to list VPC indexes for security analysis: %v", err)
	}

	for _, vpcIndex := range vpcIndexes {
		if vpcIndex == nil {
			continue
		}
		if err := client.analyzeVPCSecurity(vpcIndex); err != nil {
			// Continue with other VPCs even if one fails
			continue
		}
	}

	fmt.Println("Completed VPC security post-processing")
	return nil
}

// analyzeVPCSecurity performs comprehensive security analysis for a single VPC
func (client *boltClient) analyzeVPCSecurity(vpcIndex *types.VPCIndex) error {
	if vpcIndex == nil {
		return fmt.Errorf("vpcIndex is nil")
	}

	// Initialize security risks structure
	securityRisks := &types.VPCSecurityRisks{
		CriticalInstanceIds:        []string{},
		CriticalSecurityGroupIds:   []string{},
		CriticalLbIds:              []string{},
		HighRiskInstanceIds:        []string{},
		HighRiskSecurityGroupIds:   []string{},
		HighRiskSubnetIds:          []string{},
		MediumRiskInstanceIds:      []string{},
		MediumRiskSecurityGroupIds: []string{},
		MediumRiskAclIds:           []string{},
		UntaggedInstanceIds:        []string{},
		UntaggedSecurityGroupIds:   []string{},
		UntaggedSubnetIds:          []string{},
		UntaggedLbIds:              []string{},
		IsolatedSubnetIds:          []string{},
		OverExposedSubnetIds:       []string{},
		LastRiskAnalysis:           &time.Time{},
	}

	// Analyze instances - skip if no instances are synced
	if len(vpcIndex.InstanceIds) > 0 {
		if err := client.analyzeInstancesSecurity(vpcIndex, securityRisks); err != nil {
			return fmt.Errorf("failed to analyze instances security: %v", err)
		}
	}

	// Analyze security groups - skip if no security groups are synced
	if len(vpcIndex.SecurityGroupIds) > 0 {
		if err := client.analyzeSecurityGroupsSecurity(vpcIndex, securityRisks); err != nil {
			return fmt.Errorf("failed to analyze security groups security: %v", err)
		}
	}

	// Analyze load balancers - skip if no load balancers are synced
	if len(vpcIndex.LbIds) > 0 {
		if err := client.analyzeLoadBalancersSecurity(vpcIndex, securityRisks); err != nil {
			return fmt.Errorf("failed to analyze load balancers security: %v", err)
		}
	}

	// Analyze subnets - skip if no subnets are synced
	if len(vpcIndex.SubnetIds) > 0 {
		if err := client.analyzeSubnetsSecurity(vpcIndex, securityRisks); err != nil {
			return fmt.Errorf("failed to analyze subnets security: %v", err)
		}
	}

	// Set analysis timestamp
	now := time.Now()
	securityRisks.LastRiskAnalysis = &now

	// Update VPC Index with security risks
	vpcIndex.SecurityRisks = securityRisks
	if err := client.PutVPCIndex(vpcIndex); err != nil {
		return fmt.Errorf("failed to update VPC index with security risks: %v", err)
	}

	return nil
}

// analyzeInstancesSecurity analyzes security for all instances in the VPC
func (client *boltClient) analyzeInstancesSecurity(vpcIndex *types.VPCIndex, securityRisks *types.VPCSecurityRisks) error {
	for _, instanceId := range vpcIndex.InstanceIds {

		// Construct full CloudID for GetInstance call using provider from VPC index
		fullCloudId := fmt.Sprintf("%s:%s", vpcIndex.Provider, instanceId)

		// Get the instance details
		instance, err := client.GetInstance(fullCloudId)
		if err != nil {
			continue
		}
		if instance == nil {
			continue
		}

		// Analyze instance security
		securityStatus := client.analyzeInstanceSecurity(instance)

		// Update instance with security status
		instance.SecurityStatus = securityStatus

		if err := client.PutInstance(instance); err != nil {
			continue
		}

		// Categorize instance based on risk level
		switch securityStatus.RiskLevel {
		case types.RiskLevelCritical:
			securityRisks.CriticalInstanceIds = append(securityRisks.CriticalInstanceIds, instanceId)
		case types.RiskLevelHigh:
			securityRisks.HighRiskInstanceIds = append(securityRisks.HighRiskInstanceIds, instanceId)
		case types.RiskLevelMedium:
			securityRisks.MediumRiskInstanceIds = append(securityRisks.MediumRiskInstanceIds, instanceId)
		}

		// Check for untagged instances
		if securityStatus.IsUntagged {
			securityRisks.UntaggedInstanceIds = append(securityRisks.UntaggedInstanceIds, instanceId)
		}
	}

	return nil
}

// analyzeInstanceSecurity performs security analysis for a single instance
func (client *boltClient) analyzeInstanceSecurity(instance *types.Instance) *types.InstanceSecurityStatus {
	status := &types.InstanceSecurityStatus{
		RiskLevel:             types.RiskLevelSecure,
		IsPubliclyAccessible:  false,
		HasOpenSSHAccess:      false,
		HasOpenRDPAccess:      false,
		HasOpenDatabasePorts:  false,
		IsUntagged:            false,
		HasOverlyPermissiveSg: false,
		RiskSummary:           []string{},
		ExposedPorts:          []string{},
		RiskySecurityGroups:   []string{},
		Recommendations:       []string{},
		LastSecurityScan:      time.Now().Format(time.RFC3339),
	}

	// Check if instance has public IP
	if instance.PublicIP != "" {
		status.IsPubliclyAccessible = true
		status.Recommendations = append(status.Recommendations, "Consider placing instance in private subnet if public access is not required")
	}

	// Check if instance is untagged
	if len(instance.Labels) == 0 {
		status.IsUntagged = true
		status.RiskLevel = types.RiskLevelMedium
		status.Recommendations = append(status.Recommendations, "Add required tags for compliance and resource management")
	}

	// Analyze security groups - handle cases where instance references SGs that aren't synced
	riskyPorts := []string{}
	if instance.SecurityGroupIDs != nil {
		for _, sgId := range instance.SecurityGroupIDs {
			sg, err := client.GetSecurityGroup(sgId)
			if err != nil || sg == nil {
				// Security group not found or not synced - skip but log the missing reference
				continue
			}

			// Analyze security group rules
			risks := client.analyzeSecurityGroupRules(sg.Rules)
			if len(risks) > 0 {
				status.RiskySecurityGroups = append(status.RiskySecurityGroups, sgId)
				riskyPorts = append(riskyPorts, risks...)
			}
		}
	}

	// Check for specific risky ports
	for _, port := range riskyPorts {
		switch port {
		case "22":
			status.HasOpenSSHAccess = true
			if status.IsPubliclyAccessible {
				status.RiskLevel = types.RiskLevelCritical
				status.RiskSummary = append(status.RiskSummary, "SSH access from internet")
				status.Recommendations = append(status.Recommendations, "SSH access from internet detected - restrict to specific IP ranges or use bastion host")
			} else {
				status.RiskSummary = append(status.RiskSummary, "SSH access available")
				status.Recommendations = append(status.Recommendations, "Consider using key-based authentication and disable password authentication")
			}
		case "3389":
			status.HasOpenRDPAccess = true
			if status.IsPubliclyAccessible {
				status.RiskLevel = types.RiskLevelCritical
				status.RiskSummary = append(status.RiskSummary, "RDP access from internet")
				status.Recommendations = append(status.Recommendations, "RDP access from internet detected - restrict to specific IP ranges or use VPN")
			} else {
				status.RiskSummary = append(status.RiskSummary, "RDP access available")
				status.Recommendations = append(status.Recommendations, "Enable Network Level Authentication and strong password policies")
			}
		case "23":
			// Telnet is always critical
			status.RiskLevel = types.RiskLevelCritical
			status.RiskSummary = append(status.RiskSummary, "Telnet protocol exposed")
			status.Recommendations = append(status.Recommendations, "Disable Telnet immediately - use SSH instead")
		case "80":
			status.RiskSummary = append(status.RiskSummary, "HTTP traffic allowed")
			if status.IsPubliclyAccessible {
				status.Recommendations = append(status.Recommendations, "Consider redirecting HTTP to HTTPS for security")
			}
		case "443":
			status.RiskSummary = append(status.RiskSummary, "HTTPS traffic allowed")
			if status.IsPubliclyAccessible {
				status.Recommendations = append(status.Recommendations, "Ensure SSL certificates are up to date and use strong cipher suites")
			}
		case "3306", "5432", "1433", "1521", "27017":
			status.HasOpenDatabasePorts = true
			if status.IsPubliclyAccessible {
				status.RiskLevel = types.RiskLevelCritical
				status.RiskSummary = append(status.RiskSummary, "Database ports exposed to internet")
				status.Recommendations = append(status.Recommendations, "Database ports exposed to internet - restrict access immediately and use private subnets")
			} else {
				if status.RiskLevel == types.RiskLevelSecure {
					status.RiskLevel = types.RiskLevelMedium
				}
				status.RiskSummary = append(status.RiskSummary, "Database ports accessible within network")
				status.Recommendations = append(status.Recommendations, "Ensure database access is restricted to application servers only")
			}
		case "6379":
			// Redis - often unsecured by default
			if status.IsPubliclyAccessible {
				status.RiskLevel = types.RiskLevelCritical
				status.RiskSummary = append(status.RiskSummary, "Redis exposed to internet")
				status.Recommendations = append(status.Recommendations, "Redis exposed to internet - secure immediately with authentication and private networks")
			} else {
				if status.RiskLevel == types.RiskLevelSecure {
					status.RiskLevel = types.RiskLevelMedium
				}
				status.RiskSummary = append(status.RiskSummary, "Redis accessible within network")
				status.Recommendations = append(status.Recommendations, "Ensure Redis has authentication enabled and restrict access")
			}
		case "21":
			// FTP - insecure protocol
			status.RiskLevel = types.RiskLevelHigh
			status.RiskSummary = append(status.RiskSummary, "FTP protocol exposed")
			status.Recommendations = append(status.Recommendations, "Replace FTP with SFTP or FTPS for secure file transfer")
		case "25", "587", "465":
			// Email ports
			status.RiskSummary = append(status.RiskSummary, "Email services exposed")
			if status.IsPubliclyAccessible {
				status.Recommendations = append(status.Recommendations, "Ensure email services have proper authentication and anti-spam measures")
			}
		case "53":
			// DNS
			if status.IsPubliclyAccessible {
				status.RiskSummary = append(status.RiskSummary, "DNS service exposed to internet")
				status.Recommendations = append(status.Recommendations, "Review DNS configuration to prevent DNS amplification attacks")
			}
		case "135", "139", "445":
			// Windows file sharing
			if status.IsPubliclyAccessible {
				status.RiskLevel = types.RiskLevelCritical
				status.RiskSummary = append(status.RiskSummary, "Windows file sharing exposed to internet")
				status.Recommendations = append(status.Recommendations, "Windows file sharing exposed to internet - block immediately")
			} else {
				if status.RiskLevel == types.RiskLevelSecure {
					status.RiskLevel = types.RiskLevelMedium
				}
				status.RiskSummary = append(status.RiskSummary, "Windows file sharing accessible")
				status.Recommendations = append(status.Recommendations, "Restrict Windows file sharing to necessary systems only")
			}
		}
	}

	// Set exposed ports
	status.ExposedPorts = removeDuplicates(riskyPorts)

	// Check for unrestricted access (0.0.0.0/0 rules) - only if SG IDs exist
	if instance.SecurityGroupIDs != nil && len(instance.SecurityGroupIDs) > 0 {
		if client.hasUnrestrictedAccess(instance.SecurityGroupIDs) {
			status.HasOverlyPermissiveSg = true
			if status.RiskLevel == types.RiskLevelSecure {
				status.RiskLevel = types.RiskLevelHigh
			}
			status.Recommendations = append(status.Recommendations, "Restrict security group rules to specific IP ranges instead of 0.0.0.0/0")
		}
	}

	return status
}

// analyzeSecurityGroupsSecurity analyzes security for all security groups in the VPC
func (client *boltClient) analyzeSecurityGroupsSecurity(vpcIndex *types.VPCIndex, securityRisks *types.VPCSecurityRisks) error {
	for _, sgId := range vpcIndex.SecurityGroupIds {

		// Construct full CloudID for GetSecurityGroup call using provider from VPC index
		fullCloudId := fmt.Sprintf("%s:%s", vpcIndex.Provider, sgId)

		// Get the security group details
		sg, err := client.GetSecurityGroup(fullCloudId)
		if err != nil {
			continue
		}
		if sg == nil {
			continue
		}

		// Analyze security group security
		securityStatus := client.analyzeSecurityGroupSecurity(sg)

		// Update security group with security status
		sg.SecurityStatus = securityStatus
		if err := client.PutSecurityGroup(sg); err != nil {
			continue
		}

		// Categorize security group based on risk level
		switch securityStatus.RiskLevel {
		case types.RiskLevelCritical:
			securityRisks.CriticalSecurityGroupIds = append(securityRisks.CriticalSecurityGroupIds, sgId)
		case types.RiskLevelHigh:
			securityRisks.HighRiskSecurityGroupIds = append(securityRisks.HighRiskSecurityGroupIds, sgId)
		case types.RiskLevelMedium:
			securityRisks.MediumRiskSecurityGroupIds = append(securityRisks.MediumRiskSecurityGroupIds, sgId)
		}

		// Check for untagged security groups
		if securityStatus.IsUntagged {
			securityRisks.UntaggedSecurityGroupIds = append(securityRisks.UntaggedSecurityGroupIds, sgId)
		}
	}

	return nil
}

// analyzeSecurityGroupSecurity performs security analysis for a single security group
func (client *boltClient) analyzeSecurityGroupSecurity(sg *types.SecurityGroup) *types.SecurityGroupRiskStatus {
	status := &types.SecurityGroupRiskStatus{
		RiskLevel:                types.RiskLevelSecure,
		AllowsInternetAccess:     false,
		AllowsSSHFromInternet:    false,
		AllowsRDPFromInternet:    false,
		AllowsHTTPFromInternet:   false,
		AllowsHTTPSFromInternet:  false,
		HasOverlyPermissiveRules: false,
		IsUntagged:               false,
		RiskyRules:               []string{},
		AffectedInstances:        int32(len(sg.Instances)),
		Recommendations:          []string{},
		LastSecurityScan:         time.Now().Format(time.RFC3339),
	}

	// Check if security group is untagged
	if len(sg.Labels) == 0 {
		status.IsUntagged = true
		status.RiskLevel = types.RiskLevelMedium
		status.Recommendations = append(status.Recommendations, "Add required tags for compliance and resource management")
	}

	// Analyze rules - handle nil or empty rules safely
	if sg.Rules != nil {
		for _, rule := range sg.Rules {
			// Check for internet access (0.0.0.0/0)
			hasInternetAccess := false
			if rule.Source != nil {
				for _, source := range rule.Source {
					if strings.Contains(source, "0.0.0.0/0") || strings.Contains(source, "::/0") {
						hasInternetAccess = true
						status.AllowsInternetAccess = true
						break
					}
				}
			}

			if hasInternetAccess {
				// Check specific ports for different risk levels
				if strings.Contains(rule.PortRange, "22") {
					status.AllowsSSHFromInternet = true
					status.RiskLevel = types.RiskLevelCritical
					status.RiskyRules = append(status.RiskyRules, "SSH (22) open to internet")
					status.Recommendations = append(status.Recommendations, "Restrict SSH access to specific IP ranges or use bastion host")
				}

				if strings.Contains(rule.PortRange, "3389") {
					status.AllowsRDPFromInternet = true
					status.RiskLevel = types.RiskLevelCritical
					status.RiskyRules = append(status.RiskyRules, "RDP (3389) open to internet")
					status.Recommendations = append(status.Recommendations, "Restrict RDP access to specific IP ranges or use VPN")
				}

				if strings.Contains(rule.PortRange, "23") {
					status.RiskLevel = types.RiskLevelCritical
					status.RiskyRules = append(status.RiskyRules, "Telnet (23) open to internet")
					status.Recommendations = append(status.Recommendations, "Disable Telnet immediately - use SSH instead")
				}

				// Database ports
				dbPorts := []string{"3306", "5432", "1433", "1521", "27017"}
				for _, dbPort := range dbPorts {
					if strings.Contains(rule.PortRange, dbPort) {
						status.RiskLevel = types.RiskLevelCritical
						status.RiskyRules = append(status.RiskyRules, fmt.Sprintf("Database port %s open to internet", dbPort))
						status.Recommendations = append(status.Recommendations, "Move database to private subnet and restrict access")
					}
				}

				// Windows file sharing
				winPorts := []string{"135", "139", "445"}
				for _, winPort := range winPorts {
					if strings.Contains(rule.PortRange, winPort) {
						status.RiskLevel = types.RiskLevelCritical
						status.RiskyRules = append(status.RiskyRules, fmt.Sprintf("Windows file sharing port %s open to internet", winPort))
						status.Recommendations = append(status.Recommendations, "Block Windows file sharing from internet access")
					}
				}

				if strings.Contains(rule.PortRange, "21") {
					status.RiskLevel = types.RiskLevelHigh
					status.RiskyRules = append(status.RiskyRules, "FTP (21) open to internet")
					status.Recommendations = append(status.Recommendations, "Replace FTP with SFTP or FTPS")
				}

				if strings.Contains(rule.PortRange, "6379") {
					status.RiskLevel = types.RiskLevelCritical
					status.RiskyRules = append(status.RiskyRules, "Redis (6379) open to internet")
					status.Recommendations = append(status.Recommendations, "Secure Redis with authentication and move to private network")
				}

				if strings.Contains(rule.PortRange, "80") {
					status.AllowsHTTPFromInternet = true
					status.RiskyRules = append(status.RiskyRules, "HTTP (80) open to internet")
					if status.RiskLevel == types.RiskLevelSecure {
						status.RiskLevel = types.RiskLevelMedium
					}
					status.Recommendations = append(status.Recommendations, "Consider redirecting HTTP to HTTPS")
				}

				if strings.Contains(rule.PortRange, "443") {
					status.AllowsHTTPSFromInternet = true
					status.RiskyRules = append(status.RiskyRules, "HTTPS (443) open to internet")
					status.Recommendations = append(status.Recommendations, "Ensure SSL certificates are current and use strong ciphers")
				}

				// Check for wide port ranges
				if strings.Contains(rule.PortRange, "-") {
					parts := strings.Split(rule.PortRange, "-")
					if len(parts) == 2 {
						start, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
						end, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
						portCount := end - start + 1

						if portCount > 1000 { // Very wide range
							status.HasOverlyPermissiveRules = true
							status.RiskLevel = types.RiskLevelCritical
							status.RiskyRules = append(status.RiskyRules, fmt.Sprintf("Extremely wide port range %s (%d ports) open to internet", rule.PortRange, portCount))
							status.Recommendations = append(status.Recommendations, "Restrict to specific ports only - wide ranges are dangerous")
						} else if portCount > 100 { // Wide range
							status.HasOverlyPermissiveRules = true
							if status.RiskLevel == types.RiskLevelSecure {
								status.RiskLevel = types.RiskLevelHigh
							}
							status.RiskyRules = append(status.RiskyRules, fmt.Sprintf("Wide port range %s (%d ports) open to internet", rule.PortRange, portCount))
							status.Recommendations = append(status.Recommendations, "Narrow down port ranges to only necessary ports")
						} else if portCount > 20 { // Moderately wide range
							status.HasOverlyPermissiveRules = true
							if status.RiskLevel == types.RiskLevelSecure {
								status.RiskLevel = types.RiskLevelMedium
							}
							status.RiskyRules = append(status.RiskyRules, fmt.Sprintf("Port range %s (%d ports) open to internet", rule.PortRange, portCount))
							status.Recommendations = append(status.Recommendations, "Consider reducing port range to minimum required ports")
						}
					}
				}

				// Check for "all ports" indicators
				if strings.Contains(rule.PortRange, "0-65535") || strings.Contains(rule.PortRange, "1-65535") {
					status.HasOverlyPermissiveRules = true
					status.RiskLevel = types.RiskLevelCritical
					status.RiskyRules = append(status.RiskyRules, "All ports (0-65535) open to internet")
					status.Recommendations = append(status.Recommendations, "CRITICAL: All ports open to internet - restrict immediately")
				}
			}
		}
	}

	return status
}

// analyzeLoadBalancersSecurity analyzes security for all load balancers in the VPC
func (client *boltClient) analyzeLoadBalancersSecurity(vpcIndex *types.VPCIndex, securityRisks *types.VPCSecurityRisks) error {
	for _, lbId := range vpcIndex.LbIds {
		// Get the load balancer details
		lb, err := client.GetLB(lbId)
		if err != nil {
			fmt.Printf("WARN: Failed to get load balancer %s for security analysis: %v\n", lbId, err)
			continue
		}
		if lb == nil {
			continue
		}

		// Analyze load balancer security
		securityStatus := client.analyzeLoadBalancerSecurity(lb)

		// Update load balancer with security status
		lb.SecurityStatus = securityStatus
		if err := client.PutLB(lb); err != nil {
			fmt.Printf("WARN: Failed to update load balancer %s with security status: %v\n", lbId, err)
		}

		// Categorize load balancer based on risk level
		switch securityStatus.RiskLevel {
		case types.RiskLevelCritical:
			securityRisks.CriticalLbIds = append(securityRisks.CriticalLbIds, lbId)
		case types.RiskLevelHigh, types.RiskLevelMedium:
			// Add to high risk if medium or high for simplicity
			securityRisks.CriticalLbIds = append(securityRisks.CriticalLbIds, lbId)
		}

		// Check for untagged load balancers
		if securityStatus.IsUntagged {
			securityRisks.UntaggedLbIds = append(securityRisks.UntaggedLbIds, lbId)
		}
	}

	return nil
}

// analyzeLoadBalancerSecurity performs security analysis for a single load balancer
func (client *boltClient) analyzeLoadBalancerSecurity(lb *types.LB) *types.LoadBalancerSecurityStatus {
	status := &types.LoadBalancerSecurityStatus{
		RiskLevel:             types.RiskLevelSecure,
		IsInternetFacing:      false,
		HasInsecureListeners:  false,
		HasOpenSecurityGroups: false,
		AccessLogsDisabled:    false,
		IsUntagged:            false,
		InsecureProtocols:     []string{},
		SecurityGroupRisks:    []string{},
		BackendInstanceRisks:  []string{},
		Recommendations:       []string{},
		LastSecurityScan:      time.Now().Format(time.RFC3339),
	}

	// Check if load balancer is internet-facing
	if strings.ToLower(lb.Scheme) == "internet-facing" {
		status.IsInternetFacing = true
		status.Recommendations = append(status.Recommendations, "Ensure internet-facing load balancer is intended and properly secured")
	}

	// Check if LB is untagged
	if len(lb.Labels) == 0 {
		status.IsUntagged = true
		status.RiskLevel = types.RiskLevelMedium
		status.Recommendations = append(status.Recommendations, "Add required tags for compliance and resource management")
	}

	// Check access logs
	if !lb.AccessLogsEnabled {
		status.AccessLogsDisabled = true
		if status.RiskLevel == types.RiskLevelSecure {
			status.RiskLevel = types.RiskLevelLow
		}
		status.Recommendations = append(status.Recommendations, "Enable access logs for security monitoring and compliance")
	}

	// Analyze listeners for insecure protocols
	for _, listener := range lb.Listeners {
		if strings.ToUpper(listener.Protocol) == "HTTP" {
			status.HasInsecureListeners = true
			status.InsecureProtocols = append(status.InsecureProtocols, fmt.Sprintf("HTTP on port %d", listener.Port))
			if status.RiskLevel == types.RiskLevelSecure {
				status.RiskLevel = types.RiskLevelMedium
			}
			status.Recommendations = append(status.Recommendations, "Consider using HTTPS instead of HTTP for secure communication")
		}
	}

	// Analyze security groups
	for _, sgId := range lb.SecurityGroupIDs {
		if client.hasUnrestrictedAccess([]string{sgId}) {
			status.HasOpenSecurityGroups = true
			status.SecurityGroupRisks = append(status.SecurityGroupRisks, sgId)
			if status.RiskLevel == types.RiskLevelSecure || status.RiskLevel == types.RiskLevelLow {
				status.RiskLevel = types.RiskLevelMedium
			}
		}
	}

	// Analyze backend instances for security risks
	for _, instanceId := range lb.InstanceIDs {
		instance, err := client.GetInstance(instanceId)
		if err != nil || instance == nil {
			continue
		}

		// Check if backend instance has security issues
		if instance.SecurityStatus != nil {
			if instance.SecurityStatus.RiskLevel == types.RiskLevelCritical ||
				instance.SecurityStatus.RiskLevel == types.RiskLevelHigh {
				status.BackendInstanceRisks = append(status.BackendInstanceRisks, instanceId)
				if status.RiskLevel == types.RiskLevelSecure {
					status.RiskLevel = types.RiskLevelMedium
				}
			}
		}
	}

	if len(status.BackendInstanceRisks) > 0 {
		status.Recommendations = append(status.Recommendations, "Review and secure backend instances with security risks")
	}

	return status
}

// analyzeSubnetsSecurity analyzes security for all subnets in the VPC
func (client *boltClient) analyzeSubnetsSecurity(vpcIndex *types.VPCIndex, securityRisks *types.VPCSecurityRisks) error {
	for _, subnetId := range vpcIndex.SubnetIds {
		// Get the subnet details
		subnet, err := client.GetSubnet(subnetId)
		if err != nil {
			fmt.Printf("WARN: Failed to get subnet %s for security analysis: %v\n", subnetId, err)
			continue
		}
		if subnet == nil {
			continue
		}

		// Basic subnet analysis - check for untagged subnets
		if len(subnet.Labels) == 0 {
			securityRisks.UntaggedSubnetIds = append(securityRisks.UntaggedSubnetIds, subnetId)
		}

		// TODO: Implement more sophisticated subnet analysis
		// - Check route tables for internet gateways
		// - Identify isolated subnets
		// - Identify over-exposed subnets
	}

	return nil
}

// Helper functions

// analyzeSecurityGroupRules returns a list of risky ports from security group rules
func (client *boltClient) analyzeSecurityGroupRules(rules []types.SecurityGroupRule) []string {
	riskyPorts := []string{}

	if rules == nil {
		return riskyPorts
	}

	for _, rule := range rules {
		// Check if rule allows access from internet
		hasInternetAccess := false
		if rule.Source != nil {
			for _, source := range rule.Source {
				if strings.Contains(source, "0.0.0.0/0") || strings.Contains(source, "::/0") {
					hasInternetAccess = true
					break
				}
			}
		}

		if hasInternetAccess {
			// Extract ports from port range
			if rule.PortRange != "" {
				ports := extractPorts(rule.PortRange)
				riskyPorts = append(riskyPorts, ports...)
			}
		}
	}

	return riskyPorts
}

// hasUnrestrictedAccess checks if any security group has unrestricted (0.0.0.0/0) access
func (client *boltClient) hasUnrestrictedAccess(sgIds []string) bool {
	if sgIds == nil {
		return false
	}

	for _, sgId := range sgIds {
		sg, err := client.GetSecurityGroup(sgId)
		if err != nil || sg == nil {
			continue
		}

		if sg.Rules != nil {
			for _, rule := range sg.Rules {
				if rule.Source != nil {
					for _, source := range rule.Source {
						if strings.Contains(source, "0.0.0.0/0") || strings.Contains(source, "::/0") {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// extractPorts extracts individual ports from a port range string
func extractPorts(portRange string) []string {
	ports := []string{}

	// Handle single port
	if !strings.Contains(portRange, "-") {
		ports = append(ports, strings.TrimSpace(portRange))
		return ports
	}

	// Handle port range
	parts := strings.Split(portRange, "-")
	if len(parts) == 2 {
		start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

		if err1 == nil && err2 == nil {
			// For analysis, we'll focus on commonly known risky ports
			riskyPorts := []int{
				21,    // FTP
				22,    // SSH
				23,    // Telnet
				25,    // SMTP
				53,    // DNS
				80,    // HTTP
				135,   // Windows RPC
				139,   // NetBIOS
				443,   // HTTPS
				445,   // SMB
				465,   // SMTPS
				587,   // SMTP submission
				993,   // IMAPS
				995,   // POP3S
				1433,  // SQL Server
				1521,  // Oracle
				3306,  // MySQL
				3389,  // RDP
				5432,  // PostgreSQL
				6379,  // Redis
				27017, // MongoDB
			}
			for _, port := range riskyPorts {
				if port >= start && port <= end {
					ports = append(ports, strconv.Itoa(port))
				}
			}
		}
	}

	return ports
}

// removeDuplicates removes duplicate strings from a slice
func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}
