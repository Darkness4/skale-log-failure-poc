#!/bin/sh

go get -d -t -u ./... \
  github.com/cockroachdb/pebble@v0.0.0-20230906160148-46873a6a7a06 \
  github.com/ethereum/go-ethereum@v1.12.2 \
  github.com/crate-crypto/go-kzg-4844@v0.3.0
go mod tidy
