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
	AccountType          = "Account"
	RegionType           = "Region"
	VPCType              = "VPC"
	InstanceType         = "Instance"
	SubnetType           = "Subnet"
	ACLType              = "ACL"
	SecurityGroupType    = "SecurityGroup"
	RouteTableType       = "RouteTable"
	NATGatewayType       = "NATGateway"
	RouterType           = "Router"
	IGWType              = "IGW"
	VPCEndpointType      = "VPCEndpoint"
	PublicIPType         = "PublicIP"
	ClusterType          = "Cluster"
	PodsType             = "Pod"
	K8sServiceType       = "K8sService"
	K8sNodeType          = "K8sNode"
	NamespaceType        = "Namespace"
	LBType               = "LB"
	NetworkInterfaceType = "NetworkInterface"
	KeyPairType          = "KeyPair"
	VPNConcentratorType  = "VPNConcentrator"
)

type Error struct {
	code     int32
	message  string
	severity string
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
	ID           string
	Region       string
	VPCID        string
	PublicIP     string
	InstanceId   string
	Provider     string
	AccountID    string
	Type         string //Elastic(Static) or Dynamic
	PrivateIP    string
	Labels       map[string]string
	SelfLink     string
	LastSyncTime string
}

func (v PublicIP) DbId() string {
	return CloudID(v.Provider, v.ID)
}

func (v *PublicIP) SetSyncTime(time string) {
	v.LastSyncTime = time
}

func (v *PublicIP) GetProvider() string {
	return v.Provider
}

type Subnet struct {
	Name         string
	SubnetId     string
	CidrBlock    string
	VpcId        string
	Zone         string
	Labels       map[string]string
	Region       string
	Provider     string
	AccountID    string
	SelfLink     string
	LastSyncTime string
	RouteTableID []string
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
	SelfLink     string
	LastSyncTime string
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
	Name        string
	Destination string
	Status      string
	Target      string
	NextHopType string
	NextHopIP   string
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
	ID            string            `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	Provider      string            `json:"provider,omitempty"`
	AccountID     string            `json:"account_id,omitempty"`
	VPCId         string            `json:"vpc_id,omitempty"` //
	Region        string            `json:"region,omitempty"` // VPC Region
	State         string            `json:"state,omitempty"`
	RouteTableIds string            `json:"route_table_ids,omitempty"`
	SubnetIds     string            `json:"subnet_ids,omitempty"`
	ServiceName   string            `json:"service_name,omitempty"`
	Type          string            `json:"type,omitempty"`
	Labels        map[string]string `json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt     *time.Time        `json:"created_at,omitempty"`
	UpdatedAt     *time.Time        `json:"updated_at,omitempty"`
	SelfLink      string
	LastSyncTime  string `json:"last_sync_time,omitempty"`
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
	Name         string
	ID           string
	Provider     string
	VpcID        string
	Region       string
	Labels       map[string]string
	AccountID    string
	Rules        []SecurityGroupRule
	SelfLink     string
	LastSyncTime string
	Instances    []string
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
	TargetGroupIDs         []string
	Listeners              []LBListener
	CrossZoneLoadBalancing bool
	AccessLogsEnabled      bool
	LoggingBucket          string
	IPAddresses            []string
	IPAddressType          string
	Region                 string
	Zone                   string
	Labels                 map[string]string
	Project                string
	AccountID              string
	LastSyncTime           string
	SelfLink               string
	CreatedAt              time.Time
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

/* End resource types */
