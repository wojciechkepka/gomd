#!/bin/bash

compile() {
    local version="$1"
    export GOOS="$2"
    export GOARCH="$3"
    echo "Compiling $version-$GOOS-$GOARCH"
    go build -v -o build/gomd .
    cd build
    tar -zcvf gomd-$version-$GOOS-$GOARCH.tgz gomd
    rm gomd
    cd ..
}

command -v pkger
if [ $? != 0 ]
then
    echo "pkger required to build the binary."
    echo "download it with:"
    echo -e "\tgo get github.com/markbates/pkger/cmd/pkger"
    exit 1
fi

if [ -z $1 ]
then
    echo "USAGE:"
    echo -e "\tbuild.sh <version>"
    exit 1
fi

VERSION="$1"

rm -rf build

make pack

compile $VERSION "linux" "amd64"
compile $VERSION "darwin" "amd64"
compile $VERSION "windows" "amd64"
