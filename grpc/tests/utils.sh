#!/bin/bash


call_endpoint() {
    server="localhost:50052"
    endpoint="$1"
    response=$(grpc_cli call $server $endpoint "provider: 'aws'" --timeout=30 --json_output 2>error_file)


    # Check for endpoint errors
    if cat error_file | grep -q "Method name not found"; then
        echo "Error calling $endpoint endpoint"
        rm error_file
        exit 1
    fi
    if cat error_file | grep -q "Rpc failed with status code"; then
        error_message=$(cat error_file | sed -n '/error message:/s/.*error message: //p')
        echo "Error calling $endpoint endpoint:" $error_message
        rm error_file
        exit 1
    fi
    rm error_file


    # Check if response not empty
    if echo $response | jq -e '. == {}'  > /dev/null; then
        echo "Response from $endpoint is empty"
        exit 1
    fi

    echo $response
}

check_fields() {
    json="$1"
    required_fields="$2"

    for field in "${required_fields[@]}"; do
        field_present=$(echo "$json" | jq " has(\"$field\")")
        if [[ "$field_present" != "true" ]]; then
            echo -e "\t[x] $field is missing"
            missing_fields=$((missing_fields+1))
        fi
    done
}


