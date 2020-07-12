#/bin/bash 
echo "Building {project}..." && \
echo "Running tests..." && \
go test -v ./... && \
echo "Building Docker image..." && \
go build && \
docker build .
