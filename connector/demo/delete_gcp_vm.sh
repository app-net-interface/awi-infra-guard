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

if [ "$#" -ne 2 ]; then
    echo "This script deletes GCP Virtual Machine after it is no longer"
    echo "needed."
    echo ""
    echo "Usage: $0 RESOURCE_ID SUBREGION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for"
    echo "  creating previous resources"
    echo "SUBREGION - zone where VM will be hosted"
    echo ""
    echo "Example:"
    echo "$0 test1-1 us-east4-c"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_SUBREGION=$2

SCRIPT_VM_NAME="$SCRIPT_RES_ID-vm"

delete_vm() {
    gcloud compute firewall-rules delete demo-csp-allow-ssh-$SCRIPT_RES_ID --quiet

    echo "Attempting to delete VM: $SCRIPT_VM_NAME in zone: $SCRIPT_SUBREGION"
    gcloud compute instances delete "$SCRIPT_VM_NAME" --zone="$SCRIPT_SUBREGION" --quiet
}

delete_vm

if [ $? -eq 0 ]; then
    echo "VM successfully deleted."
else
    echo "Failed to delete VM."
    exit 1
fi

exit 0
