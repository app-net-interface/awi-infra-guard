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

if [ "$#" -ne 4 ]; then
    echo "This script creates necessary GCP Resources for testing"
    echo "creation of connection between GCP and some other provider."
    echo ""
    echo "Usage: $0 RESOURCE_ID CIDR ASN REGION"
    echo ""
    echo "RESOURE_ID - unique identifier that will be added"
    echo "  to each resource created by this script"
    echo "CIDR - unique CIDR number for subnet"
    echo "ASN - unique ASN number"
    echo "REGION - region where resources will be created"
    echo ""
    echo "Example:"
    echo "$0 test1 10.1.1.0/24 65534 us-east4"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_CIDR=$2
SCRIPT_ASN_NUMBER=$3
SCRIPT_REGION=$4

gcloud compute networks create $SCRIPT_RES_ID-vpc \
    --subnet-mode custom \
    --bgp-routing-mode global

gcloud compute networks subnets create $SCRIPT_RES_ID-subnet  \
    --network $SCRIPT_RES_ID-vpc \
    --region $SCRIPT_REGION \
    --range $SCRIPT_CIDR

gcloud compute vpn-gateways create $SCRIPT_RES_ID-vpngw \
    --network $SCRIPT_RES_ID-vpc \
    --region $SCRIPT_REGION

gcloud compute routers create $SCRIPT_RES_ID-router \
    --region $SCRIPT_REGION \
    --network $SCRIPT_RES_ID-vpc \
    --asn $SCRIPT_ASN_NUMBER \
    --advertisement-mode custom \
    --set-advertisement-groups all_subnets

exit 0
