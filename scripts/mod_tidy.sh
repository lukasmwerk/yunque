#!/bin/bash

cd "$(dirname "$0")/.."

for service in services/*; do
    if [ -d "$service" ] && [ -f "$service/go.mod" ]; then
        cd $service
        # Check if valid go dir
        if ls *.go 1> /dev/null 2>&1; then
            go mod tidy
        else
            echo "Skipping $service: No Go files found"
        fi
        cd - >/dev/null
    fi
done
