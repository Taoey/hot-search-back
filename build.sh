#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/iris-cil-server.go
chmod +x iris-cil-server

rm -rf target

mkdir -p target/web target/configs
cp -r  configs/* target/configs
cp -r web/* target/web
cp iris-cil-server target

tar -cvf iris-cil-server.tar target
rm -rf iris-cil-server
rm -rf target

