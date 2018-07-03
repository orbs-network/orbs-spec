rm -rf ./go/*
cp -r ../interfaces/* ./go
membufc --go --mock `find . -name "*.proto"`
rm `find . -name "*.proto"`