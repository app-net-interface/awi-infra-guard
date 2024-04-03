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
    echo "This script deletes GCP Resources created by other script."
    echo ""
    echo "Usage: $0 RESOURCE_ID REGION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for creating"
    echo "  these resources"
    echo "REGION - region where resources are stored"
    echo ""
    echo "Example:"
    echo "$0 test1 us-east4"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_REGION=$2

gcloud compute routers delete $SCRIPT_RES_ID-router --region $SCRIPT_REGION --quiet

gcloud compute vpn-gateways delete $SCRIPT_RES_ID-vpngw --region $SCRIPT_REGION --quiet

gcloud compute networks subnets delete $SCRIPT_RES_ID-subnet --region $SCRIPT_REGION --quiet

gcloud compute networks delete $SCRIPT_RES_ID-vpc --quiet

exit 0
