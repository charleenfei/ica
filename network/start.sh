#!/bin/bash

BINARY=icad
CHAIN_DIR=./data

echo "Starting $CHAINID in $CHAIN_DIR..."
$BINARY start --log_level warn --log_format json --home $CHAIN_DIR/$CHAINID --pruning=nothing --grpc.address="0.0.0.0:$GRPCPORT" --grpc-web.address="0.0.0.0:$GRPCWEB"
