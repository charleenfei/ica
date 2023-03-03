#!/bin/bash

set -x
set -e

make docker-reset
# make docker-build
make docker-init-chain
make docker-start-chain
make docker-init-relayer
make docker-start-relayer
make docker-init-oracle
make docker-start-oracle
make docker-unitest
