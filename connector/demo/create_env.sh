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
    echo "This script combines all other scripts to create"
    echo "quickly an environment with VMs on AWS and GCP."
    echo "It uses us-west-1 on AWS and us-east4 on GCP"
    echo ""
    echo "Usage: $0 OWNER GCP_RES_NAME AWS_RES_NAME AZURE_RES_NAME"
    echo ""
    echo "OWNER - a name of the owner creating virtual"
    echo "  machines. Each Virtual-Machine needs to be"
    echo "  tagged with the name of a person creating"
    echo "  machines."
    echo "GCP_RES_NAME - unique identifier that will be added"
    echo "  to each GCP resource created by this script"
    echo "AWS_RES_NAME - unique identifier that will be added"
    echo "  to each AWS resource created by this script"
    echo "AZURE_RES_NAME - unique identifier that will be added"
    echo "  to each Azure resource created by this script"
    echo ""
    echo "Example:"
    echo "$0 ml-training ml-data az-test"
    exit 1
fi

SCRIPT_VM_OWNER=$1
SCRIPT_GCP_RES_ID=$2
SCRIPT_AWS_RES_ID=$3
SCRIPT_AZURE_RES_ID=$4
SCRIPT_PATH="$(dirname $0)"

SCRIPT_GCP_REGION="us-east4"
SCRIPT_GCP_ZONE="us-east4-c"
SCRIPT_AWS_REGION="us-west-1"
SCRIPT_AZURE_LOCATION="westus2"

SCRIPT_GCP_ASN="$((64513 + RANDOM % 500))"
SCRIPT_GCP_CIDR="10.$((100 + RANDOM % 40)).$((RANDOM % 80)).0/24"

SCRIPT_AWS_PREFIX_CIDR="10.$((RANDOM % 40))"
SCRIPT_AWS_VPC_CIDR="$SCRIPT_AWS_PREFIX_CIDR.0.0/16"
SCRIPT_AWS_SUBNET_CIDR="$SCRIPT_AWS_PREFIX_CIDR.$((RANDOM % 240)).0/24"

SCRIPT_AZURE_ASN="$((65013 + RANDOM % 500))"
SCRIPT_AZURE_PREFIX_CIDR="10.$((50 + RANDOM % 40))"
SCRIPT_AZURE_NETWORK_CIDR="$SCRIPT_AZURE_PREFIX_CIDR.0.0/16"
SCRIPT_AZURE_GW_SUBNET_CIDR="$SCRIPT_AZURE_PREFIX_CIDR.$((RANDOM % 10)).0/24"
SCRIPT_AZURE_VM_SUBNET_CIDR="$SCRIPT_AZURE_PREFIX_CIDR.$((10 + RANDOM % 40)).0/24"

GCP_SVC_ACC="$(gcloud config get account)"
[[ "$GCP_SVC_ACC" == "" ]] && { echo "Script cannot find out the GCP Service account"; exit 1; }

echo "Creating Gateway for AWS"
set -x
$SCRIPT_PATH/create_aws_gateway.sh \
    $SCRIPT_AWS_RES_ID \
    $SCRIPT_AWS_VPC_CIDR \
    $SCRIPT_AWS_SUBNET_CIDR \
    $SCRIPT_AWS_REGION || \
        { echo "failed to create AWS Gateway"; exit 1; }
set +x

echo "Creating Gateway for GCP"
set -x
$SCRIPT_PATH/create_gcp_gateway.sh \
    $SCRIPT_GCP_RES_ID \
    $SCRIPT_GCP_CIDR \
    $SCRIPT_GCP_ASN \
    $SCRIPT_GCP_REGION || \
        { echo "failed to create GCP Gateway"; exit 1; }
set +x

echo "Creating Gateway for Azure"
set -x
$SCRIPT_PATH/create_azure_gateway.sh \
    $SCRIPT_AZURE_RES_ID \
    $SCRIPT_AZURE_NETWORK_CIDR \
    $SCRIPT_AZURE_GW_SUBNET_CIDR \
    $SCRIPT_AZURE_LOCATION \
    $SCRIPT_AZURE_ASN || \
        { echo "failed to create Azure Gateway"; exit 1; }
set +x

echo "Creating VM for AWS"
set -x
$SCRIPT_PATH/create_aws_vm.sh \
    $SCRIPT_AWS_RES_ID \
    $SCRIPT_AWS_REGION \
    $SCRIPT_VM_OWNER || \
        { echo "failed to create AWS VM"; exit 1; }
set +x

echo "Creating VM for GCP"
set -x
$SCRIPT_PATH/create_gcp_vm.sh \
    $SCRIPT_GCP_RES_ID \
    $SCRIPT_GCP_ZONE \
    $GCP_SVC_ACC \
    $SCRIPT_VM_OWNER || \
        { echo "failed to create GCP VM"; exit 1; }
set +x

echo "Creating VM for Azure"
set -x
$SCRIPT_PATH/create_azure_vm.sh \
    $SCRIPT_AZURE_RES_ID \
    $SCRIPT_AZURE_VM_SUBNET_CIDR \
    $SCRIPT_AZURE_LOCATION \
    $SCRIPT_VM_OWNER || \
        { echo "failed to create Azure VM"; exit 1; }
set +x

echo "Created successfully"
exit 0
