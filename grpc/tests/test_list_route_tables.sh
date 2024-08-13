#!/bin/bash

echo "\n-------------Test ListRouteTables-------------\n"

server="localhost:50052"

response=$(grpc_cli call $server ListRouteTables "provider: 'aws'" --json_output 2>/dev/null)

if echo $response | jq -e '. == {}'  > /dev/null; then
    echo "No route tables were found"
    exit 1
fi

# Test route tables names
test_names=(
    "ani-test-route-table"
    )

error_count=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.routeTables[]; .name == $tag)' > /dev/null; then
        echo "[V] Route table with name $test_name found"
    else
        echo "[X] Route table with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$error_count" -gt 0 ]]; then
    echo "Route tables not found: $error_count"
    exit 1
fi