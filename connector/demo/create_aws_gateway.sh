# Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
# All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http:www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

if [ "$#" -ne 4 ]; then
    echo "This script creates necessary AWS Resources for testing"
    echo "creation of connection between AWS and some other provider."
    echo ""
    echo "Usage: $0 RESOURCE_ID CIDR SUBCIDR REGION"
    echo ""
    echo "RESOURE_ID - unique identifier that will be added"
    echo "  to each resource created by this script"
    echo "CIDR - unique CIDR number for VPC"
    echo "SUBCIDR - unique CIDR number for subnet within VPC CIDR"
    echo "REGION - region where resources should be created"
    echo ""
    echo "Example:"
    echo "$0 test1 10.0.0.0/16 10.0.1.0/24 us-west-2"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_CIDR=$2
SCRIPT_SUBCIDR=$3
SCRIPT_REGION=$4

VPC_ID="$(aws ec2 describe-vpcs \
    --filters "Name=tag:Name,Values=$SCRIPT_RES_ID-vpc" \
    --query "Vpcs[].VpcId" \
    --region $SCRIPT_REGION \
    --output text)"

[[ "$VPC_ID" == "" ]] && aws ec2 create-vpc \
    --cidr-block $SCRIPT_CIDR \
    --tag-specifications "ResourceType=vpc,Tags=[{Key=Name,Value=$SCRIPT_RES_ID-vpc}]" \
    --region $SCRIPT_REGION

VPC_ID="$(aws ec2 describe-vpcs \
    --filters "Name=tag:Name,Values=$SCRIPT_RES_ID-vpc" \
    --query "Vpcs[].VpcId" \
    --region $SCRIPT_REGION \
    --output text)"

SUBNET_ID="$(aws ec2 describe-subnets \
    --filters "Name=tag:Name,Values=$SCRIPT_RES_ID-subnet" \
    --query "Subnets[].SubnetId" \
    --region $SCRIPT_REGION \
    --output text)"

[[ "$SUBNET_ID" == "" ]] && aws ec2 create-subnet \
    --vpc-id $VPC_ID \
    --tag-specifications "ResourceType=subnet,Tags=[{Key=Name,Value=$SCRIPT_RES_ID-subnet}]" \
    --cidr-block $SCRIPT_SUBCIDR \
    --region $SCRIPT_REGION

TGW_ID="$(aws ec2 describe-transit-gateways \
    --filters "Name=tag:Name,Values=AWI-DEV-TGW" "Name=state,Values=available,pending" \
    --query "TransitGateways[].TransitGatewayId" \
    --region $SCRIPT_REGION \
    --output text)"

# Missing route propagation!!!!!!
[[ "$TGW_ID" == "" ]] && aws ec2 create-transit-gateway \
    --description "My Transit Gateway" \
    --tag-specifications 'ResourceType=transit-gateway,Tags=[{Key=Name,Value=AWI-DEV-TGW}]' \
    --region $SCRIPT_REGION
# Add tag transit_route to RT


ROUTE_TABLE_ID=$( \
    aws ec2 describe-route-tables \
        --filters "Name=vpc-id,Values=$VPC_ID" \
        --region "$SCRIPT_REGION" \
        --query 'RouteTables[0].RouteTableId' \
        --output text)

[[ "$ROUTE_TABLE_ID" == "" || "$ROUTE_TABLE_ID" == "None" ]] && { echo "cannot find matching route table"; exit 1; }
[[ "$ROUTE_TABLE_ID" == *" "* ]] && { echo "found more than one matching route table. What to do: $ROUTE_TABLE_ID"; exit 1; }

aws ec2 create-tags \
    --resources "$ROUTE_TABLE_ID" \
    --region "$SCRIPT_REGION" \
    --tags "Key=transit_route,Value=True"


exit 0
