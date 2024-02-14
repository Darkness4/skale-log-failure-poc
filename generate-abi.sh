#!/bin/sh

set -ex

PROJECTPATH=$(dirname "$(realpath "$0")")

mkdir -p "${PROJECTPATH}/abi/example"
solc --evm-version paris --optimize --optimize-runs=200 "example.sol" \
  --base-path . \
  --combined-json abi,bin >"example.json"
abigen --pkg "exampleabi" \
  --combined-json "example.json" \
  --out "${PROJECTPATH}/abi/example/example.go"
rm "example.json"
