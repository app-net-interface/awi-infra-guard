#!/bin/bash

timeout=60
interval=5

start_time=$(date +%s)

while true; do
    if grpc_cli ls localhost:50052 2>/dev/null | grep -q "infra.CloudProviderService"; then
        echo "gRPC server ready"
        exit 0
    else
        echo "Waiting for gRPC server"
    fi

    current_time=$(date +%s)
    elapsed_time=$((current_time - start_time))
    if [[ $elapsed_time -ge $timeout ]]; then
        echo "gRPC server not responding - timeout" >&2
        exit 1
    fi

    sleep $interval
done

