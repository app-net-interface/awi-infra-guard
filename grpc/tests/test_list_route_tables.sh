#!/bin/bash

source ./utils.sh

echo -e "\n-------------Test ListRouteTables-------------\n"

response=$(call_endpoint ListRouteTables)
if [ $? -eq 1 ]; then
    echo $response
    exit 1
fi

# Test route tables names
test_names=(
    "ani-test-route-table"
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
    if echo "$response" | jq -e --arg tag "$test_name" 'any(.routeTables[]; .name == $tag)' > /dev/null; then
        echo "[V] Route table with name $test_name found"
        route_table=$(echo $response | jq ".routeTables[] | select(.name == \"$test_name\" )")
        check_fields "$route_table" "$required_fields"
    else
        echo "[X] Route table with name $test_name not found"
        error_count=$((error_count+1))
    fi
done

if [[ "$missing_instances" -gt 0 ]]; then
    echo "Route tables not found: $missing_instances"
    exit 1
fi

if [[ "$missing_fields" -gt 0 ]]; then
    exit 1
fi