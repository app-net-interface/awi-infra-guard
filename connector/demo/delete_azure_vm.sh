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

if [ "$#" -ne 1 ]; then
    echo "This script deletes Azure Virtual Machines "
    echo "created by other script for testing purposes."
    echo ""
    echo "Usage: $0 RESOURCE_ID"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for creating"
    echo "  these resources"
    echo ""
    echo "Example:"
    echo "$0 test1"
    exit 1
fi

SCRIPT_RES_ID=$1

SCRIPT_RES_GROUP="$SCRIPT_RES_ID-res-grp"
SCRIPT_VM_NAME="$SCRIPT_RES_ID-vm"
SCRIPT_VNET_NAME="$SCRIPT_RES_ID-vnet"
SCRIPT_SUBNET_NAME="$SCRIPT_RES_ID-subnet"
SCRIPT_NSG_NAME="$SCRIPT_RES_ID-nsg"

az vm delete \
    --resource-group "$SCRIPT_RES_GROUP" \
    --name "$SCRIPT_VM_NAME" \
    --yes

az network nsg rule delete \
    --nsg-name "$SCRIPT_NSG_NAME" \
    --name AllowSSH \
    --resource-group "$SCRIPT_RES_GROUP"

az network nsg delete \
    --name "$SCRIPT_NSG_NAME" \
    --resource-group "$SCRIPT_RES_GROUP"

SCRIPT_IP_CONFIG="$(
    az network nic list \
        --query \
            "[? \
                not_null(ipConfigurations[0].subnet) && \
                ends_with(ipConfigurations[0].subnet.id, '$SCRIPT_SUBNET_NAME') \
            ].name | [0]" \
        -o tsv)"

[[ "$SCRIPT_IP_CONFIG" == "" ]] && echo "Cannot find IP Configuration"

az network nic delete --resource-group "$SCRIPT_RES_GROUP" --name "$SCRIPT_IP_CONFIG"

az network vnet subnet delete \
    --name "$SCRIPT_SUBNET_NAME" \
    --resource-group "$SCRIPT_RES_GROUP" \
    --vnet-name "$SCRIPT_VNET_NAME"
