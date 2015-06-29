#!/bin/sh
#
# This script builds the docker image
#
cp /etc/ssl/certs/ca-certificates.crt .
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags -w -o build/castcloud
docker build -t castcloud/api .