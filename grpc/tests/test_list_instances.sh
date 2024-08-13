#!/bin/bash

echo "\n-------------Test ListInstances-------------\n"

server="localhost:50052"

response=$(grpc_cli call $server ListInstances "provider: 'aws'" --timeout=30 --json_output 2>error_file)

# Check for endpoint errors
if cat error_file | grep -q "Rpc failed with status code"; then
    error_message=$(cat error_file | sed -n '/error message:/s/.*error message: //p')
    echo "Error calling ListInstances endpoint:" $error_message
    rm error_file
    exit 1
fi
rm error_file

# Check if response not empty
if echo $response | jq -e '. == {}'  > /dev/null; then
    echo "No instances were found"
    exit 1
fi

# Test instances names
test_names=(
    "ani-test-web-server"
    )

error_count=0
for test_name in "${test_names[@]}"; do
    if echo "$response" | jq -e --arg test_name "$test_name" 'any(.instances[]; .name == $test_name and .state == "running")' > /dev/null; then
        echo "[V] Instance with name $test_name found"
    else
        echo "[X] Instance with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$error_count" -gt 0 ]]; then
    echo "Instances not found: $error_count"
    exit 1
fi