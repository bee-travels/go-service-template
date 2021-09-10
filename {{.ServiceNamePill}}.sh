#!/bin/bash

if [ "$1" = "start" ]; then
    go env -w GOBIN=$(pwd)/bin
    go install
    bin/{{.ServiceNamePill}}
elif [ "$1" = "test" ]; then
    go test
fi
