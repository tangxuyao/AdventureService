#!/usr/bin/env bash

mkdir $GOPATH/src/proto/adventure
protoc --go_out=plugins=micro:$GOPATH/src/proto/adventure ./adventure.api.proto