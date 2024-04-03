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
    echo "This script deletes AWS Resources created by other script."
    echo ""
    echo "Usage: $0 RESOURCE_ID REGION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for creating"
    echo "  these resources"
    echo "REGION - region where these resources are stored"
    echo ""
    echo "Example:"
    echo "$0 test1 us-west-2"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_REGION=$2

TGW_ID=$(aws ec2 describe-transit-gateways \
    --filters "Name=tag:Name,Values=AWI-DEV-TGW" "Name=state,Values=available,pending" \
    --query "TransitGateways[].TransitGatewayId" \
    --region $SCRIPT_REGION \
    --output text)

VPC_ID=$(aws ec2 describe-vpcs \
    --filters "Name=tag:Name,Values=$SCRIPT_RES_ID-vpc" \
    --query "Vpcs[].VpcId" \
    --region $SCRIPT_REGION \
    --output text)

[[ "$TGW_ID" != "" && "$VPC_ID" != "" ]] && TGW_VPC_ATTACHMENTS="$( \
    aws ec2 describe-transit-gateway-attachments \
        --filters \
            "Name=transit-gateway-id,Values=$TGW_ID" \
            "Name=resource-id,Values=$VPC_ID" \
            "Name=resource-type,Values=vpc" \
        --query "TransitGatewayAttachments[*].TransitGatewayAttachmentId" \
        --region $SCRIPT_REGION \
        --output text)"

[[ "$TGW_VPC_ATTACHMENTS" != "" ]] && for att_id in $TGW_VPC_ATTACHMENTS; do
    aws ec2 delete-transit-gateway-vpc-attachment \
        --transit-gateway-attachment-id $att_id \
        --region $SCRIPT_REGION || \
            { echo "failed to delete TGW with ID $att_id"; exit 1; }
done

# Wait up to 60 seconds till attachments are deleted
TIMEOUT=60

[[ "$TGW_VPC_ATTACHMENTS" != "" ]] && while true; do
    TGW_VPC_ATTACHMENTS="$( \
        aws ec2 describe-transit-gateway-attachments \
            --filters \
                "Name=transit-gateway-id,Values=$TGW_ID" \
                "Name=resource-id,Values=$VPC_ID" \
                "Name=resource-type,Values=vpc" \
            --query "TransitGatewayAttachments[*].TransitGatewayAttachmentId" \
            --region $SCRIPT_REGION \
            --output text)"
    
    [[ "$TGW_VPC_ATTACHMENTS" == "" ]] && break
    
    CURRENT_TIME=$(date +%s)
    ELAPSED_TIME=$((CURRENT_TIME - START_TIME))
    
    if [ "$ELAPSED_TIME" -ge "$TIMEOUT" ]; then
        echo "Timeout reached without an empty response. Attachments were not deleted"
        exit 1
    fi
done

[[ "$TGW_ID" != "" ]] && aws ec2 delete-transit-gateway --transit-gateway-id $TGW_ID --region $SCRIPT_REGION

SUBNET_ID=$(aws ec2 describe-subnets \
    --filters "Name=tag:Name,Values=$SCRIPT_RES_ID-subnet" \
    --query "Subnets[].SubnetId" \
    --region $SCRIPT_REGION \
    --output text)

[[ "$SUBNET_ID" != "" ]] && aws ec2 delete-subnet --subnet-id $SUBNET_ID --region $SCRIPT_REGION

[[ "$VPC_ID" != "" ]] && aws ec2 delete-vpc --vpc-id $VPC_ID --region $SCRIPT_REGION

exit 0
