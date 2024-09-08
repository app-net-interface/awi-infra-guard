#!/bin/bash

source ./utils.sh

echo -e "\n-------------Test ListVPC-------------\n"

response=$(call_endpoint ListVPC)
if [ $? -eq 1 ]; then
    echo $response
    exit 1
fi
# Test vpcs names
test_names=(
    "ani-test-vpc"
    )

# Fields required in resources
required_fields=(
    "provider"
    "accountId"
    "id"
    "region"
)

missing_fields=0
missing_instances=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.vpcs[]; .name == $tag)' > /dev/null; then
        echo "[V] VPC with name $test_name found"
        vpc=$(echo $response | jq ".vpcs[] | select(.name == \"$test_name\" )")
        check_fields "$vpc" "$required_fields"
    else
        echo "[X] VPC with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$missing_instances" -gt 0 ]]; then
    echo "VPCs not found: $missing_instances"
    exit 1
fi

if [[ "$missing_fields" -gt 0 ]]; then
    exit 1
fi