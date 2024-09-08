#!/bin/bash

source ./utils.sh

echo -e "\n-------------Test ListSubnets-------------\n"

response=$(call_endpoint ListSubnets)
if [ $? -eq 1 ]; then
    echo $response
    exit 1
fi

# Test subnets names
test_names=(
    "ani-test-subnet"
    )

# Fields required in resources
required_fields=(
    "provider"
    "accountId"
    "id"
    "region"
    "vpcId"
)

missing_fields=0
missing_instances=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.subnets[]; .name == $tag)' > /dev/null; then
        echo "[V] Subnet with name $test_name found"
        subnet=$(echo $response | jq ".subnets[] | select(.name == \"$test_name\" )")
        check_fields "$subnet" "$required_fields"
    else
        echo "[X] Subnet with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$missing_instances" -gt 0 ]]; then
    echo "Subnets not found: $missing_instances"
    exit 1
fi

if [[ "$missing_fields" -gt 0 ]]; then
    exit 1
fi