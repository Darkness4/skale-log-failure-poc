#!/bin/bash

set -euo pipefail

go run main.go \
  -rpc https://testnet.skalenodes.com/v1/aware-fake-trim-testnet \
  -ws wss://testnet.skalenodes.com/v1/ws/aware-fake-trim-testnet \
  -pk "$PRIVATE_KEY"

# go run main.go \
#   -rpc https://testnet.deepsquare.run/rpc \
#   -ws wss://testnet.deepsquare.run/ws \
#   -pk "$PRIVATE_KEY" \
