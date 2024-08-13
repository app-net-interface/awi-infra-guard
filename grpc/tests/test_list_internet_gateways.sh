#!/bin/bash

echo "\n-------------Test ListInternetGateways-------------\n"

server="localhost:50052"

response=$(grpc_cli call $server ListInternetGateways "provider: 'aws'" --json_output 2>/dev/null)

if echo $response | jq -e '. == {}'  > /dev/null; then
    echo "No internet gateways were found"
    exit 1
fi

# Test internet gateways names
test_names=(
    "ani-test-internet-gateway"
    )

error_count=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.igws[]; .name == $tag and .state == "available")' > /dev/null; then
        echo "[V] Internet gateway with name $test_name found"
    else
        echo "[X] Internet gateway with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$error_count" -gt 0 ]]; then
    echo "Internet gateways not found: $error_count"
    exit 1
fi