#!/bin/bash

server="localhost:50052"

response=$(grpc_cli call $server ListInstances "provider: 'aws'" --json_output)

# test instances names
test_names=("ani-test-web-server")

error_count=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.instances[]; .name == $tag)' > /dev/null; then
        echo "[V] Instance with name $test_name found"
    else
        echo "[X] Instance with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$error_count" -gt 0 ]]; then
    echo "$error_count instances not found" 
    exit 1
fi