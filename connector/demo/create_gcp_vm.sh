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
    echo "This script creates GCP Virtual Machine which can be used"
    echo "for testing ping connectivity."
    echo ""
    echo "Usage: $0 RESOURCE_ID SUBREGION SERVICE_ACCOUNT"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for"
    echo "  creating previous resources"
    echo "SUBREGION - zone where VM will be hosted"
    echo "SERVICE_ACCOUNT - The service account used for setting access"
    echo "  permissions for users to that VM"
    echo "OWNER - a name of the owner creating virtual"
    echo "  machines. Each Virtual-Machine needs to be"
    echo "  tagged with the name of a person creating"
    echo "  machines."
    echo ""
    echo "Example:"
    echo "$0 test1-1 us-east4-c something@something.gserviceaccount.com iosetek"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_SUBREGION=$2
SCRIPT_SERVICE_ACCOUNT=$3
SCRIPT_VM_OWNER=$4

GCLOUD_PROJECT="$(gcloud config get-value project 2>/dev/null)"

[[ "$GCLOUD_PROJECT" == "" ]] && { echo "cannot get active project"; exit 1; }

gcloud compute instances create $SCRIPT_RES_ID-vm \
    --zone=$SCRIPT_SUBREGION \
    --project="$GCLOUD_PROJECT" \
    --machine-type=e2-micro \
    --network-interface=network-tier=PREMIUM,stack-type=IPV4_ONLY,subnet=$SCRIPT_RES_ID-subnet \
    --maintenance-policy=MIGRATE \
    --hostname=$SCRIPT_RES_ID.gcp.com \
    --provisioning-model=STANDARD \
    --service-account=$SCRIPT_SERVICE_ACCOUNT \
    --scopes=https://www.googleapis.com/auth/cloud-platform \
    --create-disk=auto-delete=yes,boot=yes,device-name=$SCRIPT_RES_ID-vm,image=projects/debian-cloud/global/images/debian-11-bullseye-v20240110,mode=rw,size=10,type=projects/$GCLOUD_PROJECT/zones/$SCRIPT_SUBREGION/diskTypes/pd-balanced \
    --no-shielded-secure-boot \
    --shielded-vtpm \
    --shielded-integrity-monitoring \
    --labels=goog-ec-src=vm_add-gcloud \
    --reservation-affinity=any

# Add tags to VM
gcloud compute instances add-labels $SCRIPT_RES_ID-vm \
    --labels=app_type=ml-training \
    --zone=$SCRIPT_SUBREGION

gcloud compute instances add-labels $SCRIPT_RES_ID-vm \
    --labels=owner=$SCRIPT_VM_OWNER \
    --zone=$SCRIPT_SUBREGION

gcloud compute instances add-labels $SCRIPT_RES_ID-vm \
    --labels=project=awi \
    --zone=$SCRIPT_SUBREGION

gcloud compute firewall-rules create demo-csp-allow-ssh-$SCRIPT_RES_ID \
  --network=$SCRIPT_RES_ID-vpc \
  --allow=tcp:22 \
  --target-service-accounts=$SCRIPT_SERVICE_ACCOUNT \
  --direction=INGRESS \
  --priority=1000 \
  --description="Allow SSH access to instances with specific service account"

exit 0
