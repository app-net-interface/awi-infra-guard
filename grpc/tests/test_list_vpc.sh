#!/bin/bash

echo "\n-------------Test ListVPC-------------\n"

server="localhost:50052"

response=$(grpc_cli call $server ListVPC "provider: 'aws'" --json_output 2>/dev/null)

if echo $response | jq -e '. == {}'  > /dev/null; then
    echo "No instances were found"
    exit 1
fi

# Test vpcs names
test_names=(
    "ani-test-vpc"
    )

error_count=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.vpcs[]; .name == $tag)' > /dev/null; then
        echo "[V] VPC with name $test_name found"
    else
        echo "[X] VPC with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$error_count" -gt 0 ]]; then
    echo "VPCs not found: $error_count"
    exit 1
fi