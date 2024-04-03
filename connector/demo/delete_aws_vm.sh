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
    echo "This script deletes AWS Virtual Machine for testing"
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
SCRIPT_VPC_NAME="$SCRIPT_RES_ID-vpc"
SCRIPT_IGW_NAME="$SCRIPT_RES_ID-igw"
SCRIPT_KP_NAME="$SCRIPT_RES_ID-keypair"

# Delete VM
SCRIPT_VM_ID="$(aws ec2 describe-instances \
    --filters \
        "Name=tag:Name,Values=$SCRIPT_VM_NAME" \
        "Name=instance-state-name,Values=running,pending,shutting-down,stopping,stopped" \
    --region $SCRIPT_REGION \
    --query "Reservations[0].Instances[0].InstanceId" \
    --output text)"

if [[ "$SCRIPT_VM_ID" == *" "* ]]; then
    echo "found more than one matching VM. What to do: $SCRIPT_VM_ID"
    exit 1;
fi

if [[ "$SCRIPT_VM_ID" == "" || "$SCRIPT_VM_ID" == "None" ]]; then
    echo "Instance not found for a provided identifier. Skipping VM Deletion."
else
    aws ec2 terminate-instances \
        --region $SCRIPT_REGION \
        --instance-ids "$SCRIPT_VM_ID" || { echo "failed to start VM deletion"; exit 1; }
    
    timeout 300 aws ec2 wait instance-terminated --instance-ids "$SCRIPT_VM_ID" --region $SCRIPT_REGION || \
        { echo "failed to delete VM"; exit 1; }
fi

# Detach from VPC and delete IGW and corresponding Route from Route Table
SCRIPT_VPC_ID="$(aws ec2 describe-vpcs \
    --filters "Name=tag:Name,Values=$SCRIPT_VPC_NAME" \
    --query "Vpcs[].VpcId" \
    --region $SCRIPT_REGION \
    --output text)"

[[ "$SCRIPT_VPC_ID" == "" || "$SCRIPT_VPC_ID" == "None" ]] && { echo "cannot find matching VPC"; exit 1; }
[[ "$SCRIPT_VPC_ID" == *" "* ]] && { echo "found more than one matching VPCs. What to do: $SCRIPT_VPC_ID"; exit 1; }

SCRIPT_IGW_ID=$( \
    aws ec2 describe-internet-gateways \
        --filters "Name=tag:Name,Values=$SCRIPT_IGW_NAME" \
        --region $SCRIPT_REGION \
        --output text \
        --query 'InternetGateways[0].InternetGatewayId')

[[ "$SCRIPT_IGW_ID" == *" "* ]] && { echo "found more than one matching IGW. What to do: $SCRIPT_IGW_ID"; exit 1; }

if [[ "$SCRIPT_IGW_ID" == "" || "$SCRIPT_IGW_ID" == "None" ]]; then
    echo "Internet Gateway not found for a provided identifier. Skipping IGW Deletion."
else
    ROUTE_TABLE_ID=$( \
    aws ec2 describe-route-tables \
        --filters "Name=vpc-id,Values=$SCRIPT_VPC_ID" \
        --region "$SCRIPT_REGION" \
        --query 'RouteTables[0].RouteTableId' \
        --output text)

    [[ "$ROUTE_TABLE_ID" == "" || "$ROUTE_TABLE_ID" == "None" ]] && \
        { echo "cannot find matching route table"; exit 1; }
    [[ "$ROUTE_TABLE_ID" == *" "* ]] && \
        { echo "found more than one matching route table. What to do: $ROUTE_TABLE_ID"; exit 1; }
    
    ROUTE_CIDR="$(aws ec2 describe-route-tables \
        --route-table-ids "$ROUTE_TABLE_ID" \
        --region "$SCRIPT_REGION" \
        --query 'RouteTables[].Routes[]' \
        --output json | jq .[] | jq ". | select(.GatewayId == \"$SCRIPT_IGW_ID\")" | \
            jq .DestinationCidrBlock -r)"
    
    if [[ "$ROUTE_CIDR" == "" || "$ROUTE_CIDR" == "None" ]]; then
        echo "Route for Internet Gateway not found. Ignoring route deletion."
    else
        aws ec2 delete-route \
            --route-table-id "$ROUTE_TABLE_ID" \
            --region "$SCRIPT_REGION" \
            --destination-cidr-block "$ROUTE_CIDR" || \
                { echo "failed to remove a route to IGW"; exit 1; }
    fi

    # We could check first if the IGW is attached to a VPC but we can simply ignore
    # an error.
    aws ec2 detach-internet-gateway \
        --region $SCRIPT_REGION \
        --internet-gateway-id "$SCRIPT_IGW_ID" \
        --vpc-id "$SCRIPT_VPC_ID"
    
    aws ec2 delete-internet-gateway \
        --region $SCRIPT_REGION \
        --internet-gateway-id "$SCRIPT_IGW_ID"
fi

aws ec2 delete-key-pair \
    --region $SCRIPT_REGION \
    --key-name $SCRIPT_RES_ID-keypair
rm -f $SCRIPT_PATH/$SCRIPT_RES_ID-keypair.pem

exit 0
