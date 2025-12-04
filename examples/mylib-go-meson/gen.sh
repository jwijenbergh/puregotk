#!/usr/bin/env bash

set -e

echo "generating go files..."
go generate

tmp=$(mktemp -d)

echo "making sure imports are correct..."
GOPATH="$tmp" goimports --srcdir . -w . || { rm -r "$tmp"; exit 1; }
rm -r "$tmp"

echo "formatting files..."
go fmt "./..."

echo "running a second pass for goimports..."
goimports -w .

echo "running go vet..."
go vet ./...
