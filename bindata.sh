#!/bin/sh
go-bindata -nomemcopy -nocompress -pkg assets -o assets/bindata.go config.default.toml