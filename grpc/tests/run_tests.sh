#!/bin/bash

# List of tests to execute
tests=(
    "test_list_instances.sh"
    "test_list_vpc.sh"
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
fi