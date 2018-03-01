#!/bin/bash

if [[ $1 == "-h" || $1 == "" ]]; then
  echo "linux darwin"
  exit 0
fi

if [[ $1 == 'linux' ]]; then
  `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/linux_amd64/cider_server cider`
fi

if [[ $1 == 'darwin' ]]; then
  `CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/darwin/cider_server cider`
fi
