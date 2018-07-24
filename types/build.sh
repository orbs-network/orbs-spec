echo "Make sure your membuffers git repo is updated and pulled before building!"
echo ""

rm -rf ./go/*
cp -r ../interfaces/* ./go

### OLD: (uses brew) membufc --go --mock `find . -name "*.proto"`
### NEW: running membufc directly from source to avoid waiting for brew releases
go run $(ls -1 ../../membuffers/go/membufc/*.go | grep -v _test.go) --go --mock `find . -name "*.proto"`
rm `find . -name "*.proto"`

echo ""
echo "Building all packages with go build:"

command 2>&1 go build -a -v ./go/... | grep "orbs-network\|.mb.go"