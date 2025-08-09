// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
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

package types

import (
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	ConditionLabel = "condition"
	OrCondition    = "OR"
	AndCondition   = "AND"
	Separator      = ":"
)

const (
	AccountType            = "Account"
	RegionType             = "Region"
	VPCType                = "VPC"
	InstanceType           = "Instance"
	SubnetType             = "Subnet"
	ACLType                = "ACL"
	SecurityGroupType      = "SecurityGroup"
	RouteTableType         = "RouteTable"
	NATGatewayType         = "NATGateway"
	RouterType             = "Router"
	IGWType                = "IGW"
	VPCEndpointType        = "VPCEndpoint"
	VPCIndexType           = "VPCIndex"
	PublicIPType           = "PublicIP"
	ClusterType            = "Cluster"
	PodsType               = "Pod"
	K8sServiceType         = "K8sService"
	K8sNodeType            = "K8sNode"
	NamespaceType          = "Namespace"
	LBType                 = "LB"
	NetworkInterfaceType   = "NetworkInterface"
	KeyPairType            = "KeyPair"
	VPNConcentratorType    = "VPNConcentrator"
	VPCConnectionType      = "VPCConnection"
	VPCConnectionGraphType = "VPCConnectionGraph"
)

// Risk level constants for security analysis
const (
	RiskLevelSecure   = "SECURE"
	RiskLevelLow      = "LOW_RISK"
	RiskLevelMedium   = "MEDIUM_RISK"
	RiskLevelHigh     = "HIGH_RISK"
	RiskLevelCritical = "CRITICAL_RISK"
)

type Error struct {
	code     int32
	message  string
	severity string
}

// Security status structures for cross-resource risk analysis

// VPCSecurityRisks represents overall security metrics for a VPC
type VPCSecurityRisks struct {
	// Critical risk resources (immediate attention required)
	CriticalInstanceIds      []string `json:"critical_instance_ids,omitempty"`       // Instances with public IP + open SSH/RDP
	CriticalSecurityGroupIds []string `json:"critical_security_group_ids,omitempty"` // SGs with 0.0.0.0/0 on critical ports
	CriticalLbIds            []string `json:"critical_lb_ids,omitempty"`             // Internet-facing LBs with risky backends
	// High risk resources (significant security concerns)
	HighRiskInstanceIds      []string `json:"high_risk_instance_ids,omitempty"`       // Public instances with database ports exposed
	HighRiskSecurityGroupIds []string `json:"high_risk_security_group_ids,omitempty"` // SGs with wide port ranges open
	HighRiskSubnetIds        []string `json:"high_risk_subnet_ids,omitempty"`         // Subnets with mixed public/private resources
	// Medium risk resources (best practice violations)
	MediumRiskInstanceIds      []string `json:"medium_risk_instance_ids,omitempty"`       // Untagged instances
	MediumRiskSecurityGroupIds []string `json:"medium_risk_security_group_ids,omitempty"` // Unused or overly permissive SGs
	MediumRiskAclIds           []string `json:"medium_risk_acl_ids,omitempty"`            // ACLs with broad rules
	// Compliance violations
	UntaggedInstanceIds      []string `json:"untagged_instance_ids,omitempty"`
	UntaggedSecurityGroupIds []string `json:"untagged_security_group_ids,omitempty"`
	UntaggedSubnetIds        []string `json:"untagged_subnet_ids,omitempty"`
	UntaggedLbIds            []string `json:"untagged_lb_ids,omitempty"`
	// Network topology risks
	IsolatedSubnetIds    []string   `json:"isolated_subnet_ids,omitempty"`     // Subnets with no route to internet
	OverExposedSubnetIds []string   `json:"over_exposed_subnet_ids,omitempty"` // Subnets with multiple IGW routes
	LastRiskAnalysis     *time.Time `json:"last_risk_analysis,omitempty"`      // When risk analysis was performed
}

// InstanceSecurityStatus represents security analysis for an instance
type InstanceSecurityStatus struct {
	RiskLevel             string   `json:"overall_risk_level,omitempty"`       // SECURE, LOW_RISK, MEDIUM_RISK, HIGH_RISK, CRITICAL_RISK
	IsPubliclyAccessible  bool     `json:"is_publicly_accessible,omitempty"`   // Quick UI flag
	HasOpenSSHAccess      bool     `json:"has_open_ssh_access,omitempty"`      // Quick UI flag
	HasOpenRDPAccess      bool     `json:"has_open_rdp_access,omitempty"`      // Quick UI flag
	HasOpenDatabasePorts  bool     `json:"has_open_database_ports,omitempty"`  // Database ports exposed to internet
	IsUntagged            bool     `json:"is_untagged,omitempty"`              // Quick UI flag
	HasOverlyPermissiveSg bool     `json:"has_overly_permissive_sg,omitempty"` // Security groups with broad access
	RiskSummary           []string `json:"risk_summary,omitempty"`             // Human-readable risk descriptions
	ExposedPorts          []string `json:"exposed_ports,omitempty"`            // List of risky ports that are exposed
	RiskySecurityGroups   []string `json:"risky_security_groups,omitempty"`    // SG IDs that contribute to risk
	Recommendations       []string `json:"recommendations,omitempty"`          // Actionable security recommendations
	LastSecurityScan      string   `json:"last_assessed,omitempty"`            // Timestamp of last analysis
}

// SecurityGroupRiskStatus represents security analysis for a security group
type SecurityGroupRiskStatus struct {
	RiskLevel                string   `json:"risk_level,omitempty"`                  // SECURE, LOW_RISK, MEDIUM_RISK, HIGH_RISK, CRITICAL_RISK
	AllowsInternetAccess     bool     `json:"allows_internet_access,omitempty"`      // Quick UI flag
	AllowsSSHFromInternet    bool     `json:"allows_ssh_from_internet,omitempty"`    // Quick UI flag
	AllowsRDPFromInternet    bool     `json:"allows_rdp_from_internet,omitempty"`    // Quick UI flag
	AllowsHTTPFromInternet   bool     `json:"allows_http_from_internet,omitempty"`   // Quick UI flag
	AllowsHTTPSFromInternet  bool     `json:"allows_https_from_internet,omitempty"`  // Quick UI flag
	HasOverlyPermissiveRules bool     `json:"has_overly_permissive_rules,omitempty"` // Quick UI flag
	IsUntagged               bool     `json:"is_untagged,omitempty"`                 // Quick UI flag
	RiskyRules               []string `json:"risky_rules,omitempty"`                 // List of concerning rules
	AffectedInstances        int32    `json:"affected_instances,omitempty"`          // Count of instances using this SG
	Recommendations          []string `json:"recommendations,omitempty"`             // Actionable security recommendations
	LastSecurityScan         string   `json:"last_security_scan,omitempty"`          // Timestamp of last analysis
}

// LoadBalancerSecurityStatus represents security analysis for a load balancer
type LoadBalancerSecurityStatus struct {
	RiskLevel             string   `json:"risk_level,omitempty"`               // SECURE, LOW_RISK, MEDIUM_RISK, HIGH_RISK, CRITICAL_RISK
	IsInternetFacing      bool     `json:"is_internet_facing,omitempty"`       // Quick UI flag
	HasInsecureListeners  bool     `json:"has_insecure_listeners,omitempty"`   // Quick UI flag for HTTP vs HTTPS
	HasOpenSecurityGroups bool     `json:"has_open_security_groups,omitempty"` // Quick UI flag
	AccessLogsDisabled    bool     `json:"access_logs_disabled,omitempty"`     // Quick UI flag
	IsUntagged            bool     `json:"is_untagged,omitempty"`              // Quick UI flag
	InsecureProtocols     []string `json:"insecure_protocols,omitempty"`       // List of HTTP listeners
	SecurityGroupRisks    []string `json:"security_group_risks,omitempty"`     // List of risky security groups
	BackendInstanceRisks  []string `json:"backend_instance_risks,omitempty"`   // List of risky backend instances
	Recommendations       []string `json:"recommendations,omitempty"`          // Actionable security recommendations
	LastSecurityScan      string   `json:"last_security_scan,omitempty"`       // Timestamp of last analysis
}

/* Start SyncTime types */

type SyncTime struct {
	Provider     string
	ResourceType string
	Time         string
}

func (v *SyncTime) DbId() string {
	return SyncTimeKey(v.Provider, v.ResourceType)
}

func (v *SyncTime) SetSyncTime(time string) {

}

func (v *SyncTime) GetProvider() string {
	return v.Provider
}

func CloudID(provider, id string) string {
	return provider + ":" + id
}

func SyncTimeKey(provider string, typ string) string {
	return fmt.Sprintf("%s/%s", provider, typ)
}

func SyncTimeKeyDecode(s string) (provider, type_ string, err error) {
	split := strings.Split(s, "/")
	if len(split) != 2 {
		return "", "", fmt.Errorf("failed to determine provider and type from key %s", s)
	}
	return split[0], split[1], nil
}

/* End SyncTime types */

/* Start Connection types */
type DestinationDetails struct {
	Provider string
	VPC      string
	Region   string
}

type SingleVPCConnectionParams struct {
	ConnID      string
	VpcID       string
	Region      string
	Destination DestinationDetails
}

type VPCConnectionParams struct {
	ConnID  string
	Vpc1ID  string
	Vpc2ID  string
	Region1 string
	Region2 string
}

type VPCConnectionOutput struct {
	Region1 string
	Region2 string
}

type SingleVPCConnectionOutput struct {
	Region string
}

type VPCDisconnectionParams struct {
	ConnID  string
	Vpc1ID  string
	Vpc2ID  string
	Region1 string
	Region2 string
}

type SingleVPCDisconnectionParams struct {
	ConnID string
	VpcID  string
	Region string
}

type VPCDisconnectionOutput struct {
}

/* End Connection types */

type Region struct {
	ID           string
	Name         string
	Provider     string
	AccountID    string
	LastSyncTime string
}

func (r *Region) DbId() string {
	return CloudID(r.Provider, r.ID)
}

func (r *Region) SetSyncTime(time string) {
	r.LastSyncTime = time
}

func (r *Region) GetProvider() string {
	return r.Provider
}

/* Start resource types */

type VPC struct {
	ID           string
	Name         string
	Region       string
	Labels       map[string]string
	IPv4CIDR     string
	IPv6CIDR     string
	Provider     string
	AccountID    string
	SelfLink     string
	Project      string
	LastSyncTime string
}

func (v *VPC) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *VPC) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *VPC) GetProvider() string {
	return v.Provider
}

type Instance struct {
	ID               string
	Name             string
	PublicIP         string
	PrivateIP        string
	SubnetID         string
	VPCID            string
	Labels           map[string]string
	State            string
	Project          string
	Region           string
	Zone             string
	Provider         string
	AccountID        string
	Type             string
	SecurityGroupIDs []string
	InterfaceIDs     []string
	LastSyncTime     string
	SelfLink         string
	Cost             float64                 `json:"cost,omitempty"` // Added field for cost
	CreatedAt        *time.Time              `json:"created_at,omitempty"`
	UpdatedAt        *time.Time              `json:"updated_at,omitempty"`
	SecurityStatus   *InstanceSecurityStatus `json:"security_status,omitempty"` // Security analysis results
}

func (v Instance) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *Instance) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Instance) GetProvider() string {
	return v.Provider
}

type PublicIP struct {
	ID                 string            `json:"id,omitempty"`
	Type               string            `json:"type,omitempty"`
	Provider           string            `json:"provider,omitempty"`
	AccountID          string            `json:"account_id,omitempty"`
	VPCId              string            `json:"vpc_id,omitempty"` //
	Region             string            `json:"region,omitempty"`
	PublicIP           string            `json:"public_ip,omitempty"`
	InstanceId         string            `json:"instance_id,omitempty"`
	PrivateIP          string            `json:"private_ip,omitempty"`
	Byoip              string            `json:"byoip,omitempty"`
	Project            string            `json:"project,omitempty"`
	SelfLink           string            `json:"self_link,omitempty"`
	Labels             map[string]string `json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt          *time.Time        `json:"created_at,omitempty"`
	UpdatedAt          *time.Time        `json:"updated_at,omitempty"`
	LastSyncTime       string            `json:"last_sync_time,omitempty"`
	NetworkInterfaceId string            `json:"network_interface_id,omitempty"` // Added field
}

func (v *PublicIP) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *PublicIP) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *PublicIP) GetProvider() string {
	return v.Provider
}

type Subnet struct {
	Name          string
	SubnetId      string
	CidrBlock     string
	VpcId         string
	Zone          string
	Labels        map[string]string
	Region        string
	Provider      string
	AccountID     string
	SelfLink      string
	RouteTableIds []string
	NetworkAclIds []string
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	LastSyncTime  string     `json:"last_sync_time,omitempty"`
}

func (v *Subnet) DbId() string {
	return CloudID(v.Provider, v.SubnetId)
}

func (v *Subnet) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Subnet) GetProvider() string {
	return v.Provider
}

type Ports []string

type ProtocolsAndPorts map[string]Ports

type Account struct {
	Name         string
	ID           string
	Provider     string
	LastSyncTime string
}

type RouteTable struct {
	Name         string
	ID           string
	Provider     string
	VpcID        string
	Region       string
	Labels       map[string]string
	AccountID    string
	Routes       []Route
	SubnetIds    []string
	IGWIds       []string
	VGWIds       []string
	NGWIds       []string
	TGWIds       []string
	SelfLink     string
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	LastSyncTime string     `json:"last_sync_time,omitempty"`
}

func (v *RouteTable) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *RouteTable) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *RouteTable) GetProvider() string {
	return v.Provider
}

type Route struct {
	Name                   string
	Destination            string
	Status                 string
	Target                 string
	NextHopType            string
	NextHopIP              string
	DestinationVpcID       string // VPC ID when this is a peering route
	VpcPeeringConnectionID string // ID of VPC peering connection if this is a peering route
}

// CloudGateway represents a generic cloud gateway with various attributes.
type Router struct {
	ID                   string            `json:"id"`
	AccountID            string            `json:"account_id,omitempty"`
	Name                 string            `json:"name"`
	Provider             string            `json:"provider"`
	Region               string            `json:"region"`
	VPCId                string            `json:"vpc_id"`
	State                string            `json:"state"`
	AdvertisedRange      string            `json:"advertised_range"`
	AdvertisedGroup      string            `json:"advertised_group"`
	SubnetId             string            `json:"subnet_id"`
	ASN                  uint32            `json:"asn"`
	CIDRBlock            string            `json:"cidr_block"`
	StaticRoutes         []string          `json:"static_routes"` // Could be a list of CIDR blocks
	VPNType              string            `json:"vpn_type"`
	SecurityGroupIDs     []string          `json:"security_group_ids"` // Security groups or ACLs IDs
	Labels               map[string]string `json:"labels"`
	VPCAttachments       []VPCAttachment   `json:"vpc_attachments"` // Attached VPCs (for transit gateways etc)
	CreatedAt            time.Time         `json:"created_at"`
	UpdatedAt            time.Time         `json:"updated_at"`
	AdditionalProperties map[string]string `json:"additional_properties"`
	SelfLink             string
	LastSyncTime         string `json:"last_sync_time"`
}

func (v *Router) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *Router) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *Router) GetProvider() string {
	return v.Provider
}

type NATGateway struct {
	ID                   string
	Name                 string            `json:"name,omitempty"`
	Provider             string            `json:"provider,omitempty"`
	AccountID            string            `json:"account_id,omitempty"`
	VpcId                string            `json:"vpc_id,omitempty"`
	Region               string            `json:"region,omitempty"`
	State                string            `json:"state,omitempty"`
	PublicIp             string            `json:"public_ip,omitempty"`
	PrivateIp            string            `json:"private_ip,omitempty"`
	SubnetId             string            `json:"subnet_id,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"`
	CreatedAt            time.Time         `json:"created_at,omitempty"`
	UpdatedAt            time.Time         `json:"updated_at,omitempty"`
	LastSyncTime         string            `json:"last_sync_time,omitempty"`
	SelfLink             string
	AdditionalProperties map[string]string `json:"additional_properties,omitempty"`
}

func (v *NATGateway) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *NATGateway) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *NATGateway) GetProvider() string {
	return v.Provider
}

type IGW struct {
	ID            string                 `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Provider      string                 `json:"provider,omitempty"`
	AccountID     string                 `json:"account_id,omitempty"`
	AttachedVpcId string                 `json:"attached_vpc_id,omitempty"` //
	Region        string                 `json:"region,omitempty"`          // VPC Region
	State         string                 `json:"state,omitempty"`
	Labels        map[string]string      `json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt     *timestamppb.Timestamp `json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `json:"updated_at,omitempty"`
	SelfLink      string
	LastSyncTime  string `json:"last_sync_time,omitempty"`
}

func (v *IGW) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *IGW) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *IGW) GetProvider() string {
	return v.Provider
}

type VPCEndpoint struct {
	ID               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Provider         string            `json:"provider,omitempty"`
	AccountID        string            `json:"account_id,omitempty"`
	VPCId            string            `json:"vpc_id,omitempty"` //
	Region           string            `json:"region,omitempty"` // VPC Region
	State            string            `json:"state,omitempty"`
	RouteTableIds    []string          `json:"route_table_ids,omitempty"`    // Changed from string to []string
	SubnetIds        []string          `json:"subnet_ids,omitempty"`         // Changed from string to []string
	SecurityGroupIDs []string          `json:"security_group_ids,omitempty"` // Added field
	ServiceName      string            `json:"service_name,omitempty"`
	Type             string            `json:"type,omitempty"`
	Labels           map[string]string `json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt        *time.Time        `json:"created_at,omitempty"`
	UpdatedAt        *time.Time        `json:"updated_at,omitempty"`
	SelfLink         string
	LastSyncTime     string `json:"last_sync_time,omitempty"`
}

func (v *VPCEndpoint) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *VPCEndpoint) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *VPCEndpoint) GetProvider() string {
	return v.Provider
}

type SecurityGroup struct {
	Name           string
	ID             string
	Provider       string
	VpcID          string
	Region         string
	Labels         map[string]string
	AccountID      string
	Rules          []SecurityGroupRule
	SelfLink       string
	LastSyncTime   string
	Instances      []string
	SecurityStatus *SecurityGroupRiskStatus `json:"security_status,omitempty"` // Security analysis results
}

type SecurityGroupRule struct {
	Protocol  string
	PortRange string
	Source    []string
	Direction string
}

func (v *SecurityGroup) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *SecurityGroup) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *SecurityGroup) GetProvider() string {
	return v.Provider
}

type Firewall struct {
	Name         string
	ID           string
	Provider     string
	VpcID        string
	Region       string
	Labels       map[string]string
	AccountID    string
	Rules        []ACLRule
	SelfLink     string
	LastSyncTime string
}

type ACL struct {
	Name         string
	ID           string
	Provider     string
	VpcID        string
	Region       string
	Labels       map[string]string
	AccountID    string
	Rules        []ACLRule
	SelfLink     string
	LastSyncTime string
	Subnets      []string
}

type ACLRule struct {
	Number            int
	Protocol          string
	PortRange         string
	SourceRanges      []string
	DestinationRanges []string
	Action            string
	Direction         string
}

func (v *ACL) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *ACL) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *ACL) GetProvider() string {
	return v.Provider
}

// LBListener represents a listener for a load balancer
type LBListener struct {
	ListenerID    string
	Protocol      string
	Port          int32
	TargetGroupID string
}

// LB represents a load balancer
type LB struct {
	ID                     string
	Name                   string
	Provider               string
	Type                   string
	DNSName                string
	Scheme                 string
	VPCID                  string
	InstanceIDs            []string
	SubnetIDs              []string
	SecurityGroupIDs       []string
	TargetGroupIDs         []string
	Listeners              []LBListener
	CrossZoneLoadBalancing bool
	AccessLogsEnabled      bool
	LoggingBucket          string
	PublicIPs              []string
	PrivateIPs             []string
	State                  string
	IPAddressType          string
	Region                 string
	Zone                   string
	Labels                 map[string]string
	Project                string
	AccountID              string
	LastSyncTime           string
	SelfLink               string
	CreatedAt              time.Time
	SecurityStatus         *LoadBalancerSecurityStatus `json:"security_status,omitempty"` // Security analysis results
}

func (lb *LB) DbId() string {
	return CloudID(lb.Provider, lb.ID)
}

func (lb *LB) SetSyncTime(t string) {
	lb.LastSyncTime = t
}

func (lb *LB) GetProvider() string {
	return lb.Provider
}

// End LBtype
type NetworkInterface struct {
	ID               string
	Name             string
	Provider         string
	AccountID        string
	VPCID            string
	InstanceID       string
	SubnetID         string
	AvailabilityZone string
	Region           string
	Labels           map[string]string
	PrivateIPs       []string
	PublicIP         string
	SecurityGroupIDs []string
	MacAddress       string
	PrivateDNSName   string
	PublicDNSName    string
	Description      string
	Status           string
	SelfLink         string
	InterfaceType    string
	LastSyncTime     string
}

func (n *NetworkInterface) DbId() string {
	return CloudID(n.Provider, n.ID)
}

func (n *NetworkInterface) SetSyncTime(time string) {
	n.LastSyncTime = time
}

func (n *NetworkInterface) GetProvider() string {
	return n.Provider
}

type KeyPair struct {
	ID           string
	Name         string
	Region       string
	Fingerprint  string
	PublicKey    string
	KeyPairType  string
	CreatedAt    time.Time
	Labels       map[string]string
	Provider     string
	AccountID    string
	LastSyncTime string
	InstanceIds  []string
}

func (k *KeyPair) DbId() string {
	return CloudID(k.Provider, k.ID)
}

func (k *KeyPair) SetSyncTime(time string) {
	k.LastSyncTime = time
}

func (k *KeyPair) GetProvider() string {
	return k.Provider
}

type VPCAttachment struct {
	VpcId     string            `json:"vpc_id"`
	SubnetIds []string          `json:"subnet_ids"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Labels    map[string]string `json:"labels"`
}

type VPNConcentrator struct {
	ID           string
	Name         string
	Provider     string
	AccountID    string
	VpcID        string
	Region       string
	State        string
	Type         string
	ASN          int64
	Labels       map[string]string
	CreatedAt    time.Time
	LastSyncTime string
	SelfLink     string
}

func (v *VPNConcentrator) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *VPNConcentrator) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *VPNConcentrator) GetProvider() string {
	return v.Provider
}

type VPCIndex struct {
	VpcId               string                 `json:"vpc_id,omitempty"`
	InstanceIds         []string               `json:"instance_ids,omitempty"`
	AclIds              []string               `json:"acl_ids,omitempty"`
	SecurityGroupIds    []string               `json:"security_group_ids,omitempty"`
	NatGatewayIds       []string               `json:"nat_gateway_ids,omitempty"`
	VpcEndpointIds      []string               `json:"vpc_endpoint_ids,omitempty"`
	LbIds               []string               `json:"lb_ids,omitempty"`
	RouterIds           []string               `json:"router_ids,omitempty"`
	IgwIds              []string               `json:"igw_ids,omitempty"`
	SubnetIds           []string               `json:"subnet_ids,omitempty"`
	RouteTableIds       []string               `json:"route_table_ids,omitempty"`
	NetworkInterfaceIds []string               `json:"network_interface_ids,omitempty"`
	KeyPairIds          []string               `json:"key_pair_ids,omitempty"`
	VpnConcentratorIds  []string               `json:"vpn_concentrator_ids,omitempty"`
	PublicIpIds         []string               `json:"public_ip_ids,omitempty"`
	ClusterIds          []string               `json:"cluster_ids,omitempty"`
	LastSyncTime        string                 `json:"last_sync_time,omitempty"`
	Provider            string                 `json:"provider,omitempty"`
	AccountId           string                 `json:"account_id,omitempty"`
	Region              string                 `json:"region,omitempty"`
	CreatedAt           *timestamppb.Timestamp `json:"created_at,omitempty"`
	UpdatedAt           *timestamppb.Timestamp `json:"updated_at,omitempty"`
	SecurityRisks       *VPCSecurityRisks      `json:"security_risks,omitempty"` // Overall VPC security metrics
}

func (v *VPCIndex) DbId() string {
	return CloudID(v.Provider, v.VpcId)
}

func (v *VPCIndex) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *VPCIndex) GetProvider() string {
	return v.Provider
}

/* End resource types */

// VpcGraphNode represents a resource in the VPC connectivity graph.
type VpcGraphNode struct {
	ID           string            `json:"id,omitempty"`            // Unique ID of the resource (e.g., instance ID, subnet ID)
	ResourceType string            `json:"resource_type,omitempty"` // Type of the resource (e.g., "Instance", "Subnet", "RouteTable")
	Name         string            `json:"name,omitempty"`          // Display name of the resource
	Properties   map[string]string `json:"properties,omitempty"`    // Key-value pairs for essential display properties (e.g., "privateIP", "cidrBlock", "state")
	Provider     string            `json:"provider,omitempty"`
	AccountId    string            `json:"account_id,omitempty"`
	Region       string            `json:"region,omitempty"`
}

// DbId returns a unique identifier for the VpcGraphNode, combining provider and resource ID.
// Note: VpcGraphNode itself might not be stored directly in a DB this way,
// but this follows the pattern if needed.
func (v *VpcGraphNode) DbId() string {
	// Use the node's own ID field, not VpcId
	return CloudID(v.Provider, v.ID)
}

// GetProvider returns the provider for the VpcGraphNode.
func (v *VpcGraphNode) GetProvider() string {
	return v.Provider
}

// VpcGraphEdge represents a connection or relationship between two nodes in the graph.
type VpcGraphEdge struct {
	SourceNodeID     string `json:"source_node_id,omitempty"`    // ID of the source node
	TargetNodeID     string `json:"target_node_id,omitempty"`    // ID of the target node
	RelationshipType string `json:"relationship_type,omitempty"` // Describes the relationship (e.g., "CONTAINS", "ROUTES_TO", "ASSOCIATED_WITH")
	Provider         string `json:"provider,omitempty"`
	AccountId        string `json:"account_id,omitempty"`
	Region           string `json:"region,omitempty"`
}

// InstanceGraphNode represents a resource in the Instance connectivity graph.
type InstanceGraphNode struct {
	ID           string            `json:"id,omitempty"`            // Unique ID of the resource
	ResourceType string            `json:"resource_type,omitempty"` // Type of the resource
	Name         string            `json:"name,omitempty"`          // Display name
	Properties   map[string]string `json:"properties,omitempty"`    // Key-value pairs for display
	Provider     string            `json:"provider,omitempty"`
	AccountID    string            `json:"account_id,omitempty"`
	Region       string            `json:"region,omitempty"`
}

// InstanceGraphEdge represents a connection in the Instance connectivity graph.
type InstanceGraphEdge struct {
	SourceNodeID     string `json:"source_node_id,omitempty"`
	TargetNodeID     string `json:"target_node_id,omitempty"`
	RelationshipType string `json:"relationship_type,omitempty"`
	Provider         string `json:"provider,omitempty"`
	AccountID        string `json:"account_id,omitempty"`
	Region           string `json:"region,omitempty"`
}

// Add any helper methods for VpcGraphNode or VpcGraphEdge if needed below

// VPCConnection represents a connection between two VPCs, potentially across different accounts
type VPCConnection struct {
	Provider         string
	ID               string
	Name             string
	Account          string // Owner/primary account of the connection
	Region           string // Primary region of the connection
	FromVpcId        string
	ToVpcId          string
	FromVpcAccountId string // Account ID for the first VPC
	ToVpcAccountId   string // Account ID for the second VPC
	FromVpcRegion    string // Region for the first VPC
	ToVpcRegion      string // Region for the second VPC
	Status           string
	ConnectionType   string // Type of connection (peering, transit_gateway, etc.)
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Tags             map[string]string // Additional metadata tags
	LastSyncTime     string            // Time of last sync operation
}

// VPCConnection helper methods to implement db.DbObject interface
func (v *VPCConnection) GetProvider() string {
	return v.Provider
}

func (v *VPCConnection) DbId() string {
	return v.ID
}

func (v *VPCConnection) SetSyncTime(t string) {
	// empty since VPC Connections don't need sync time
}

type VpcInternalGraph struct {

	// All nodes representing VPCs, Transit Gateways, etc.
	Nodes []VpcGraphNode `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// All edges representing connections between nodes
	Edges []VpcGraphEdge `protobuf:"bytes,3,rep,name=edges,proto3" json:"edges,omitempty"`
	// Cross-account/region information
	AccountId    string            `protobuf:"bytes,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Region       string            `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Provider     string            `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
	Labels       map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastSyncTime string            `protobuf:"bytes,8,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
}

type VpcConnectionGraph struct {
	Id string `json:"id,omitempty"`
	// All nodes in the graph (VPCs, Transit Gateways, VPC Endpoints, etc)
	Nodes []*VpcConnectionGraphNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// All edges representing connections between nodes
	Edges []*VpcConnectionGraphEdge `protobuf:"bytes,2,rep,name=edges,proto3" json:"edges,omitempty"`

	SrcVpcGraph  *VpcInternalGraph `protobuf:"bytes,3,opt,name=src_vpc_graph,json=srcVpcGraph,proto3" json:"src_vpc_graph,omitempty"`
	DestVpcGraph *VpcInternalGraph `protobuf:"bytes,4,opt,name=dest_vpc_graph,json=destVpcGraph,proto3" json:"dest_vpc_graph,omitempty"` // The source VPC connectivity graph this is derived from

	// Graph metadata
	Accounts []string `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"` // List of account IDs included in the graph
	Regions  []string `protobuf:"bytes,4,rep,name=regions,proto3" json:"regions,omitempty"`   // List of regions included in the graph
	Provider string   `json:"provider,omitempty"`

	Labels       map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastSyncTime string            `protobuf:"bytes,6,opt,name=last_sync_time,json=lastSyncTime,proto3" json:"last_sync_time,omitempty"`
}

type VpcConnectionGraphNode struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	NodeType string `json:"node_type,omitempty"`
	// Node properties based on type
	//
	// Types that are assignable to TypeProperties:
	//
	//	*VpcConnectionGraphNode_Vpc
	//	*VpcConnectionGraphNode_Tgw
	//	*VpcConnectionGraphNode_Endpoint
	// Cross-account/region information
	AccountId string `json:"account_id,omitempty"`
	Region    string `json:"region,omitempty"`
	Provider  string `json:"provider,omitempty"`
	// Additional metadata
	Properties map[string]string `json:"properties,omitempty"`
	Labels     map[string]string `json:"labels,omitempty"`
}

type VpcConnectionGraphEdge struct {
	Id           string `json:"id,omitempty"`
	SourceNodeId string `json:"source_node_id,omitempty"`
	TargetNodeId string `json:"target_node_id,omitempty"`
	// The type of connection this edge represents
	ConnectionType string `json:"connection_type,omitempty"`
	// Connection details based on type
	//
	// Types that are assignable to ConnectionDetails:
	//
	//	*VpcConnectionGraphEdge_PeeringDetails
	//	*VpcConnectionGraphEdge_TransitGatewayDetails
	//	*VpcConnectionGraphEdge_EndpointDetails
	//	*VpcConnectionGraphEdge_TransitVpcDetails
	// Edge properties
	Status        string `json:"status,omitempty"`        // active, pending, failed
	Bidirectional bool   `json:"bidirectional,omitempty"` // Whether traffic can flow both ways
	// Cross-account/region information
	AccountId string `json:"account_id,omitempty"`
	Region    string `json:"region,omitempty"`
	Provider  string `json:"provider,omitempty"`
	// Cross-account/region metadata
	RouteTableIds []string          `json:"route_table_ids,omitempty"`                                       // Route tables implementing this connection
	Properties    map[string]string `json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3"` // Additional properties (bandwidth, latency, etc)
}

// VPCConnection helper methods to implement db.DbObject interface
func (v *VpcConnectionGraph) GetProvider() string {
	return v.Provider
}

func (v *VpcConnectionGraph) DbId() string {
	return v.Id
}

func (v *VpcConnectionGraph) SetSyncTime(t string) {
	// empty since VPC Connections don't need sync time
}
