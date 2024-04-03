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

if [ "$#" -ne 5 ]; then
    echo "This script creates necessary Azure Resources for testing"
    echo "creation of connection between Azure and some other provider."
    echo ""
    echo "Usage: $0 RESOURCE_ID CIDR SUBCIDR ASN LOCATION"
    echo ""
    echo "RESOURE_ID - unique identifier that will be added"
    echo "  to each resource created by this script"
    echo "LOCATION - location where resources should be created"
    echo "CIDR - The CIDR prefix for created VNET"
    echo "SUBCIDR - The CIDR prefix for created subnet (inside vnet)"
    echo "Example:"
    echo "$0 test1 15.0.0.0/16 15.0.0.0/24 westus2 64876"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_CIDR=$2
SCRIPT_SUBCIDR=$3
SCRIPT_LOCATION=$4
SCRIPT_ASN=$5

SCRIPT_RES_GROUP=$SCRIPT_RES_ID-res-grp
SCRIPT_VNET_NAME=$SCRIPT_RES_ID-vnet

az group create \
    --name "$SCRIPT_RES_GROUP" \
    --location "$SCRIPT_LOCATION"

# The guide tells about disabling Firewall for VNet - explore that

az network vnet create \
    --name "$SCRIPT_VNET_NAME" \
    --resource-group "$SCRIPT_RES_GROUP" \
    --location "$SCRIPT_LOCATION" \
    --address-prefixes "$SCRIPT_CIDR" \
    --subnet-name GatewaySubnet \
    --subnet-prefix "$SCRIPT_SUBCIDR"

az network public-ip create \
    --name $SCRIPT_RES_ID-ip1 \
    --resource-group "$SCRIPT_RES_GROUP" \
    --location "$SCRIPT_LOCATION" \
    --allocation-method Static

az network public-ip create \
    --name $SCRIPT_RES_ID-ip2 \
    --resource-group "$SCRIPT_RES_GROUP" \
    --location "$SCRIPT_LOCATION" \
    --allocation-method Static

# The VPN Gateway with BGP IP Addresses defined cannot be
# used for the communication with the AWS, since AWS provides
# it's own IP Addresses for BGP.
#
# Even though VPN Gateway BGP Addresses can be updated, this
# is not something we should do - if something else already
# acknowledged these IP Addresses, it could make a mess.
az network vnet-gateway create \
    --name $SCRIPT_RES_ID-vnet-gw \
    --resource-group "$SCRIPT_RES_GROUP" \
    --location "$SCRIPT_LOCATION" \
    --vnet "$SCRIPT_VNET_NAME" \
    --gateway-type Vpn \
    --vpn-type RouteBased \
    --sku VpnGw1 \
    --asn "$SCRIPT_ASN" \
    --public-ip-addresses $SCRIPT_RES_ID-ip1 $SCRIPT_RES_ID-ip2


