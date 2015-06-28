#!/bin/sh
#
# This script generates embedded data
#
go-bindata -nomemcopy -nocompress -pkg assets -o assets/bindata.go config.default.toml