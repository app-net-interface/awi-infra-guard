#!/bin/bash

source ./utils.sh

echo -e "\n-------------Test ListInternetGateways-------------\n"

response=$(call_endpoint ListInternetGateways)
if [ $? -eq 1 ]; then
    echo $response
    exit 1
fi

# Test internet gateways names
test_names=(
    "ani-test-internet-gateway"
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
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.igws[]; .name == $tag and .state == "available")' > /dev/null; then
        echo "[V] Internet gateway with name $test_name found"
        internet_gateway=$(echo $response | jq ".igws[] | select(.name == \"$test_name\" )")
        check_fields "$internet_gateway" "$required_fields"
    else
        echo "[X] Internet gateway with name $test_name not found"
        missing_instances=$((missing_instances+1))
    fi
done

if [[ "$missing_instances" -gt 0 ]]; then
    echo "Internet gateways not found: $missing_instances"
    exit 1
fi

if [[ "$missing_fields" -gt 0 ]]; then
    exit 1
fi