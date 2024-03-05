# Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
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

if [ "$#" -ne 3 ]; then
    echo "This script combines all other scripts to clean up"
    echo "quickly an environment with VMs on AWS and GCP."
    echo "It uses us-west-1 on AWS and us-east4 on GCP"
    echo ""
    echo "Usage: $0 GCP_RES_NAME AWS_RES_NAME AZURE_RES_NAME"
    echo ""
    echo "GCP_RES_NAME - unique identifier that was used"
    echo "  for creating the GCP resources earlier"
    echo "AWS_RES_NAME - unique identifier that was used"
    echo "  for creating the AWS resources earlier"
    echo "AZURE_RES_NAME - unique identifier that was used"
    echo "  for creating the Azure resources earlier"
    echo ""
    echo "Example:"
    echo "$0 ml-training ml-data az-test"
    exit 1
fi

SCRIPT_GCP_RES_ID=$1
SCRIPT_AWS_RES_ID=$2
SCRIPT_AZURE_RES_ID=$3
SCRIPT_PATH="$(dirname $0)"

SCRIPT_GCP_REGION="us-east4"
SCRIPT_GCP_ZONE="us-east4-c"
SCRIPT_AWS_REGION="us-west-1"
SCRIPT_AZURE_REGION="westus2"

echo "Deleting AWS VMs"
set -x
$SCRIPT_PATH/delete_aws_vm.sh \
    $SCRIPT_AWS_RES_ID \
    $SCRIPT_AWS_REGION
set +x

echo "Deleting GCP VMs"
set -x
$SCRIPT_PATH/delete_gcp_vm.sh \
    $SCRIPT_GCP_RES_ID \
    $SCRIPT_GCP_ZONE && \
        $SCRIPT_PATH/wait_for_vm_deletion_gcp.sh \
            $SCRIPT_GCP_RES_ID \
            $SCRIPT_GCP_ZONE \
            300
set +x

echo "Deleting Azure VMs"
set -x
$SCRIPT_PATH/delete_azure_vm.sh \
    $SCRIPT_AWS_RES_ID
set +x

echo "Deleting AWS Gateway"
set -x
$SCRIPT_PATH/delete_aws_gateway.sh \
    $SCRIPT_AWS_RES_ID \
    $SCRIPT_AWS_REGION
set +x

echo "Deleting GCP Gateway"
set -x
$SCRIPT_PATH/delete_gcp_gateway.sh \
    $SCRIPT_GCP_RES_ID \
    $SCRIPT_GCP_REGION
set +x

echo "Deleting Azure Gateway"
set -x
$SCRIPT_PATH/delete_azure_gateway.sh \
    $SCRIPT_AZURE_RES_ID \
    $SCRIPT_AZURE_REGION
set +x

echo "Deleted successfully."
