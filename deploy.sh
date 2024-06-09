#!/bin/bash

echo "BUILDING NEW BIMATRI CONTAINER............."
docker build -t bima-tri .
if [ $? -ne 0 ]; then
    echo "Failed to build Docker image"
    exit 1
fi

echo "STOPPING EXISTING BIMATRI CONTAINER............."
docker container stop bima-tri
if [ $? -eq 0 ]; then
    echo "Container stopped successfully"
else
    echo "No running container named bima-tri, skipping stop step"
fi

echo "REMOVING EXISTING BIMATRI CONTAINER............."
docker container rm bima-tri
if [ $? -eq 0 ]; then
    echo "Container removed successfully"
else
    echo "No container named bima-tri, skipping remove step"
fi

echo "RUNNING NEW BIMATRI CONTAINER............."
docker run --restart always -d -p 8787:8787 --name bima-tri bima-tri
if [ $? -ne 0 ]; then
    echo "Failed to run new Docker container"
    exit 1
fi

echo "ALL IS DONE!"
