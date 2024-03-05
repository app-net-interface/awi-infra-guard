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
    echo "This script creates Azure Virtual Machine for testing"
    echo "ping connectivity along with resources needed for"
    echo "reaching the machine with SSH."
    echo ""
    echo "Usage: $0 RESOURCE_ID SUBNET_CIDR LOCATION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for"
    echo "  creating Azure resources"
    echo "SUBNET_CIDR - the cidr address for subnet where VMs"
    echo "  will be spawned"
    echo "LOCATION - location where gateway resources were created"
    echo ""
    echo "Example:"
    echo "$0 test1 15.0.1.0/24 westus2"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_SUBCIDR=$2
SCRIPT_LOCATION=$3

SCRIPT_PATH="$(dirname $0)"

SCRIPT_RES_GROUP="$SCRIPT_RES_ID-res-grp"
SCRIPT_VM_NAME="$SCRIPT_RES_ID-vm"
SCRIPT_VNET_NAME="$SCRIPT_RES_ID-vnet"
SCRIPT_SUBNET_NAME="$SCRIPT_RES_ID-subnet"
SCRIPT_NSG_NAME="$SCRIPT_RES_ID-nsg"

az network vnet subnet create \
    --vnet-name "$SCRIPT_VNET_NAME" \
    --resource-group "$SCRIPT_RES_GROUP" \
    --name "$SCRIPT_SUBNET_NAME" \
    --address-prefix "$SCRIPT_SUBCIDR"


az vm create \
  --resource-group "$SCRIPT_RES_GROUP" \
  --name "$SCRIPT_VM_NAME" \
  --image Ubuntu2204 \
  --size Standard_B1ls \
  --admin-username azureuser \
  --authentication-type ssh \
  --generate-ssh-keys \
  --vnet-name "$SCRIPT_VNET_NAME" \
  --subnet "$SCRIPT_SUBNET_NAME"

az network nsg create \
    --name "$SCRIPT_NSG_NAME" \
    --resource-group "$SCRIPT_RES_GROUP" \
    --location "$SCRIPT_LOCATION"

az network nsg rule create \
  --resource-group "$SCRIPT_RES_GROUP" \
  --nsg-name "$SCRIPT_NSG_NAME" \
  --name AllowSSH \
  --priority 100 \
  --direction Inbound \
  --access Allow \
  --protocol Tcp \
  --source-address-prefix '*' \
  --source-port-range '*' \
  --destination-address-prefix '*' \
  --destination-port-range 22

exit 0
