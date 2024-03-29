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
    echo "This script deletes Azure Resources created by other script."
    echo ""
    echo "Usage: $0 RESOURCE_ID LOCATION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for creating"
    echo "  these resources"
    echo "LOCATION - location where resources are stored"
    echo ""
    echo "Example:"
    echo "$0 test1 westus2"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_LOCATION=$2

SCRIPT_RES_GROUP="$SCRIPT_RES_ID-res-grp"

az network vnet-gateway delete \
    --resource-group "$SCRIPT_RES_GROUP" \
    --name $SCRIPT_RES_ID-vnet-gw

az network public-ip delete \
    --name $SCRIPT_RES_ID-ip1 \
    --resource-group $SCRIPT_RES_GROUP 

az network public-ip delete \
    --name $SCRIPT_RES_ID-ip2 \
    --resource-group $SCRIPT_RES_GROUP

az network vnet delete \
    --name $SCRIPT_RES_ID-vnet \
    --resource-group $SCRIPT_RES_GROUP

az group create \
    --name $SCRIPT_RES_GROUP \
    --location $SCRIPT_LOCATION
