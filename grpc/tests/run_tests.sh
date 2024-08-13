#!/bin/bash

# List of tests to execute
tests=(
    "test_list_instances.sh"
    "test_list_vpc.sh"
    "test_list_route_tables.sh"
    "test_list_subnets.sh"
    "test_list_internet_gateways.sh"
    )

for test in "${tests[@]}"; do
    sh $test
    if [ $? -eq 1 ]; then
        test_error=1
    fi
done

if [ $test_error -eq 1 ]; then
    echo "Tests failed"
    exit 1
else
    echo "Test passed"
fi