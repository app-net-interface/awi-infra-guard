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

if [ "$#" -ne 2 ]; then
    echo "This script creates AWS Virtual Machine for testing"
    echo "ping connectivity along with Internet Gateway/Security Rules"
    echo "etc. needed for reaching the machine with SSH."
    echo ""
    echo "Usage: $0 RESOURCE_ID REGION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for"
    echo "  creating AWS resources"
    echo "REGION - region where resources were be created"
    echo ""
    echo "Example:"
    echo "$0 test1-1 us-west-2"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_REGION=$2

SCRIPT_PATH="$(dirname $0)"

SCRIPT_VM_NAME="$SCRIPT_RES_ID-vm"

# Get VPC ID
SCRIPT_VPC_NAME="$SCRIPT_RES_ID-vpc"

SCRIPT_VPC_ID="$(aws ec2 describe-vpcs \
    --filters "Name=tag:Name,Values=$SCRIPT_VPC_NAME" \
    --query "Vpcs[].VpcId" \
    --region $SCRIPT_REGION \
    --output text)"

[[ "$SCRIPT_VPC_ID" == "" || "$SCRIPT_VPC_ID" == "None" ]] && { echo "cannot find matching VPC"; exit 1; }
[[ "$SCRIPT_VPC_ID" == *" "* ]] && { echo "found more than one matching VPCs. What to do: $SCRIPT_VPC_ID"; exit 1; }

# Get Subnet ID
SCRIPT_SUBNET_ID="$( \
    aws ec2 describe-subnets \
        --region $SCRIPT_REGION \
        --output text \
        --filters "Name=vpc-id,Values=$SCRIPT_VPC_ID" \
        --query "Subnets[*].{ID:SubnetId}" 2>/dev/null)"

[[ "$SCRIPT_SUBNET_ID" == "" || "$SCRIPT_SUBNET_ID" == "None" ]] && { echo "cannot find matching subnet"; exit 1; }
[[ "$SCRIPT_SUBNET_ID" == *" "* ]] && { echo "found more than one matching subnet. What to do: $SCRIPT_SUBNET_ID"; exit 1; }

# TODO: Check if Keypair already exists

# Create KeyPair
aws ec2 create-key-pair \
    --key-name $SCRIPT_RES_ID-keypair \
    --query 'KeyMaterial' \
    --region $SCRIPT_REGION \
    --output text > $SCRIPT_PATH/$SCRIPT_RES_ID-keypair.pem || { echo "failed to create KeyPair"; exit 1; }

chmod 400 $SCRIPT_PATH/$SCRIPT_RES_ID-keypair.pem || { echo "failed to set key permissions"; exit 1; }

# Get Default SG
SCRIPT_SG="$( \
    aws ec2 describe-security-groups \
        --region $SCRIPT_REGION \
        --filters "Name=vpc-id,Values=$SCRIPT_VPC_ID" "Name=group-name,Values=default" \
        --query 'SecurityGroups[0].GroupId' \
        --output text)"

[[ "$SCRIPT_SG" == "" || "$SCRIPT_SG" == "None" ]] && { echo "cannot find matching security group"; exit 1; }
[[ "$SCRIPT_SG" == *" "* ]] && { echo "found more than one security group. What to do: $SCRIPT_SG"; exit 1; }

# Create Internet Gateway if doesn't exist
SCRIPT_IGW_NAME=$SCRIPT_RES_ID-igw

IGW_INFO=$( \
    aws ec2 describe-internet-gateways \
        --filters "Name=tag:Name,Values=$SCRIPT_IGW_NAME" \
        --region $SCRIPT_REGION \
        --output text \
        --query 'InternetGateways')

[[ "$IGW_INFO" == "" || "$IGW_INFO" == "None" ]] && aws ec2 create-internet-gateway \
    --query 'InternetGateway.{InternetGatewayId:InternetGatewayId}' \
    --region $SCRIPT_REGION \
    --output json | jq '.InternetGatewayId' | \
        xargs -I {} aws ec2 create-tags \
            --resources {} \
            --region $SCRIPT_REGION \
            --tags Key=Name,Value=$SCRIPT_IGW_NAME

SCRIPT_IGW_ID=$( \
    aws ec2 describe-internet-gateways \
        --filters "Name=tag:Name,Values=$SCRIPT_IGW_NAME" \
        --region $SCRIPT_REGION \
        --output text \
        --query 'InternetGateways[0].InternetGatewayId')

[[ "$SCRIPT_IGW_ID" == "" || "$SCRIPT_IGW_ID" == "None" ]] && { echo "cannot find matching IGW"; exit 1; }

# Attach Internet Gateway to VPC
SCRIPT_IGW_ATTACHED_VPC=$( \
    aws ec2 describe-internet-gateways \
        --filters "Name=tag:Name,Values=$SCRIPT_IGW_NAME" \
        --region $SCRIPT_REGION \
        --output text \
        --query 'InternetGateways[0].Attachments[0].VpcId')

[[ \
    "$SCRIPT_IGW_ATTACHED_VPC" != "" && \
    "$SCRIPT_IGW_ATTACHED_VPC" != "None" && \
    "$SCRIPT_IGW_ATTACHED_VPC" != "$SCRIPT_VPC_ID" \
]] && \
    { echo "IGW is already attached to a different VPC: $SCRIPT_IGW_ATTACHED_VPC"; exit 1; }

[[ "$SCRIPT_IGW_ATTACHED_VPC" == "" || "$SCRIPT_IGW_ATTACHED_VPC" == "None" ]] && \
    aws ec2 attach-internet-gateway \
        --region "$SCRIPT_REGION" \
        --vpc-id "$SCRIPT_VPC_ID" \
        --internet-gateway-id "$SCRIPT_IGW_ID"

# Add Route Entry to push remaining traffic through IGW
ROUTE_TABLE_ID=$( \
    aws ec2 describe-route-tables \
        --filters "Name=vpc-id,Values=$SCRIPT_VPC_ID" \
        --region "$SCRIPT_REGION" \
        --query 'RouteTables[0].RouteTableId' \
        --output text)

[[ "$ROUTE_TABLE_ID" == "" || "$ROUTE_TABLE_ID" == "None" ]] && { echo "cannot find matching route table"; exit 1; }
[[ "$ROUTE_TABLE_ID" == *" "* ]] && { echo "found more than one matching route table. What to do: $ROUTE_TABLE_ID"; exit 1; }

# We assume that failure here can mean that the Route already exists which is fine
# We should probably handle it better.
aws ec2 create-route \
    --route-table-id "$ROUTE_TABLE_ID" \
    --region "$SCRIPT_REGION" \
    --destination-cidr-block "0.0.0.0/0" \
    --gateway-id "$SCRIPT_IGW_ID"

# Allow SSH From ingress
aws ec2 authorize-security-group-ingress \
    --group-id "$SCRIPT_SG" \
    --region "$SCRIPT_REGION" \
    --protocol tcp \
    --port 22 \
    --cidr "0.0.0.0/0"


# Run Instance
aws ec2 run-instances \
    --image-id ami-0353faff0d421c70e \
    --region $SCRIPT_REGION \
    --count 1 \
    --instance-type t2.nano \
    --key-name $SCRIPT_RES_ID-keypair \
    --network-interfaces "SubnetId=$SCRIPT_SUBNET_ID,AssociatePublicIpAddress=true,DeviceIndex=0,Groups=$SCRIPT_SG" \
    --tag-specifications "ResourceType=instance,Tags=[{Key=Name,Value=$SCRIPT_VM_NAME}]" \
    --metadata-options "HttpTokens=required,HttpEndpoint=enabled,HttpPutResponseHopLimit=2" \
    --private-dns-name-options "HostnameType=ip-name,EnableResourceNameDnsARecord=false,EnableResourceNameDnsAAAARecord=false"

# Add tag to VM

SCRIPT_VM_ID="$(aws ec2 describe-instances \
    --filters \
        "Name=tag:Name,Values=$SCRIPT_VM_NAME" \
    --region $SCRIPT_REGION \
    --query "Reservations[0].Instances[0].InstanceId" \
    --output text)"

aws ec2 create-tags \
    --resources $SCRIPT_VM_ID \
    --region $SCRIPT_REGION \
    --tags Key=app_type,Value=ml-data

exit 0
