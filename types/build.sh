#!/bin/sh

echo "Make sure your membuffers git repo is updated and pulled before building!"
echo ""

rm -rf ./go/*
cp -r ../interfaces/* ./go
cp -r ./additional-go/argument_array_builder ./go/protocol/argument_array_builder.go

### OLD: (uses brew) membufc --go --mock `find . -name "*.proto"`
### NEW: running membufc directly from source to avoid waiting for brew releases
### NEW2: running membufc directly via a small main that calls the api of compile.
go run compile.go `find . -name "*.proto"`
rm `find . -name "*.proto"`

echo ""
echo "Formatting all go files with go fmt:"
go fmt ./go/...

echo ""
echo "Building all packages with go build:"

command 2>&1 go build -a -v ./go/... | grep "orbs-network\|.mb.go"
