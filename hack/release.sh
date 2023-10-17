#!/usr/bin/env bash

OS="linux darwin windows"
ARCHITECTURES="amd64 arm64"

VERSION=$1
NAME=$2
MODULE_NAME=$3

for arch in $ARCHITECTURES; do
  for os in $OS; do
    echo "Building $os-$arch"
    mkdir -p "_output/$os-$arch"
    if [ "$os" == "windows" ]; then
      GOOS=$os GOARCH=$arch go build -ldflags "-s -w -X $MODULE_NAME/pkg/version.Version=$VERSION" -o _output/$os-$arch/$NAME.exe ./cmd/
    else
      GOOS=$os GOARCH=$arch go build -ldflags "-s -w -X $MODULE_NAME/pkg/version.Version=$VERSION" -o _output/$os-$arch/$NAME ./cmd/
    fi
    cd _output
    tar -zcvf $NAME-$os-$arch.tar.gz $os-$arch
    rm -rf $os-$arch
    cd ..
  done
done