#!/usr/bin/env bash

export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ~/cosmos/sbip/data/test-1) # && echo $WALLET_1;
# export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ~/cosmos/sbip/data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') # && echo $ICA_ADDR

export product="amazon.org"
export price="50"

envsubst < message.txt > message.json
#echo "Final message.json"
cat message.json | jq .
