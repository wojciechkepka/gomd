#!/bin/bash

compile() {
    local version="$1"
    export GOOS="$2"
    export GOARCH="$3"
    echo "Compiling $version-$GOOS-$GOARCH"
    go build -o build/gomd .
    cd build
    tar -zcvf gomd-$version-$GOOS-$GOARCH.tgz gomd
    rm gomd
    cd ..
}

if [ -z $1 ]
then
    echo "USAGE:"
    echo -e "\tbuild.sh <version>"
    exit 1
fi

VERSION="$1"

rm -rf build

compile $VERSION "linux" "amd64"
compile $VERSION "darwin" "amd64"
compile $VERSION "windows" "amd64"
