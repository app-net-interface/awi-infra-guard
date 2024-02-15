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
    echo "This script waits for the VM to be up with a timeout set."
    echo ""
    echo "Usage: $0 RESOURCE_ID REGION"
    echo ""
    echo "RESOURE_ID - unique identifier that was used for"
    echo "  creating previous resources"
    echo "REGION - region where VM will be hosted"
    echo "TIMEOUT - maximum time  in seconds for the VM to start."
    echo "  Returns 1 if timeout reached"
    echo ""
    echo "Example:"
    echo "$0 test1 us-west-2 300"
    exit 1
fi

SCRIPT_RES_ID=$1
SCRIPT_REGION=$2
SCRIPT_TIMEOUT=$3

SCRIPT_VM_NAME="$SCRIPT_RES_ID-vm"

SCRIPT_VM_ID="$( \
    aws ec2 describe-instances \
        --filters "Name=tag:Name,Values=$SCRIPT_VM_NAME" \
        --region "$SCRIPT_REGION" \
        --query 'Reservations[0].Instances[0].[InstanceId]' \
        --output text)"

[[ "$SCRIPT_VM_ID" == "" || "$SCRIPT_VM_ID" == "None" ]] && { echo "cannot find matching VM"; exit 1; }
[[ "$SCRIPT_VM_ID" == *" "* ]] && { echo "found more than one matching VMs. What to do: $SCRIPT_VM_ID"; exit 1; }

timeout $SCRIPT_TIMEOUT aws ec2 wait instance-running --instance-ids "$SCRIPT_VM_ID" || \
    { echo "the VM did not start"; exit 1; }

exit 0
