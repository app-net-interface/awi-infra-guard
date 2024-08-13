#!/bin/bash

echo "\n-------------Test ListSubnets-------------\n"

server="localhost:50052"

response=$(grpc_cli call $server ListSubnets "provider: 'aws'" --json_output 2>/dev/null)

if echo $response | jq -e '. == {}'  > /dev/null; then
    echo "No subnets were found"
    exit 1
fi

# Test subnets names
test_names=(
    "ani-test-subnet"
    )

error_count=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.subnets[]; .name == $tag)' > /dev/null; then
        echo "[V] Subnet with name $test_name found"
    else
        echo "[X] Subnet with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$error_count" -gt 0 ]]; then
    echo "Subnets not found: $error_count"
    exit 1
fi