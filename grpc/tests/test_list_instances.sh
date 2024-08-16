#!/bin/bash

source ./utils.sh

echo -e "\n-------------Test ListInstances-------------\n"

response=$(call_endpoint ListInstances)
if [ $? -eq 1 ]; then
    echo $response
    exit 1
fi

# Test instances names
test_names=(
    "ani-test-web-server"
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
    if echo "$response" | jq -e --arg test_name "$test_name" 'any(.instances[]; .name == $test_name and .state == "running")' > /dev/null; then
        echo "[V] Instance with name $test_name found"
        instance=$(echo $response | jq ".instances[] | select(.name == \"$test_name\" and .state == \"running\" )")
        check_fields "$instance" "$required_fields"
    else
        echo "[X] Instance with name $test_name not found"
        missing_instances=$((missing_instances+1))
    fi
done

if [[ "$missing_instances" -gt 0 ]]; then
    echo "Instances not found: $missing_instances"
    exit 1
fi

if [[ "$missing_fields" -gt 0 ]]; then
    exit 1
fi