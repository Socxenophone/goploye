#!/bin/bash

# Set the Go module to the project's path
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

# Define the project's path
PROJECT_PATH="github.com/Socxenophone/goploye"

# Build the binary with optimizations
go build -o goploye -ldflags="-s -w" $PROJECT_PATH

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful! Optimized binary is at ./goploye"
else
    echo "Build failed"
fi
