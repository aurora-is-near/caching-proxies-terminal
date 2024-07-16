#!/bin/bash

# Enable script to exit on error and print commands and their arguments as they are executed.
set -euxo pipefail

echo "Building Docker image 'caching-proxies-terminal:latest'..."
docker build --no-cache -t "caching-proxies-terminal:latest" -f Dockerfile .
if [ $? -eq 0 ]; then
    echo "Docker image built successfully."
else
    echo "Docker image build failed."
    exit 1
fi

echo "Done building caching-proxies-terminal:latest"

echo "Process completed successfully."