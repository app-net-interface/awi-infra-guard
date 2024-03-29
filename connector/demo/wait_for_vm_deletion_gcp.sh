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
    echo "This script waits for the VM to be removed with a timeout set."
    echo ""
    echo "Usage: $0 RESOURCE_ID SUBREGION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for"
    echo "  creating previous resources"
    echo "SUBREGION - zone where the VM was created"
    echo "TIMEOUT - maximum time in seconds for the VM to be deleted."
    echo "  Returns 1 if timeout reached"
    echo ""
    echo "Example:"
    echo "$0 test1 us-east4-c 300"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_SUBREGION=$2
SCRIPT_TIMEOUT=$3

SCRIPT_VM_NAME="$SCRIPT_RES_ID-vm"

vm_exists() {
    gcloud compute instances describe "$SCRIPT_VM_NAME" --zone="$SCRIPT_SUBREGION" --format="value(name)" 2>/dev/null
}

SCRIPT_START_TIME=$(date +%s)

# Main loop that waits for the VM to be in the RUNNING state
while true; do
    SCRIPT_CURRENT_TIME=$(date +%s)
    SCRIPT_ELAPSED_TIME=$(( SCRIPT_CURRENT_TIME - SCRIPT_START_TIME ))

    # Check if the elapsed time is greater than the timeout
    if [ "$SCRIPT_ELAPSED_TIME" -ge "$SCRIPT_TIMEOUT" ]; then
        echo "Timeout reached. Exiting."
        exit 1
    fi

    if ! vm_exists; then
        echo "VM successfully deleted."
        break
    else
        echo "Waiting for VM to be deleted..."
        sleep 10
    fi
done

exit 0
