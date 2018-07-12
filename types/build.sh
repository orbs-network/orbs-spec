rm -rf ./go/*
cp -r ../interfaces/* ./go
membufc --go --mock `find . -name "*.proto"`
rm `find . -name "*.proto"`
echo ""
echo "Building all packages with go build:"
command 2>&1 go build -a -v ./go/... | grep "orbs-network\|.mb.go"