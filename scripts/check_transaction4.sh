#!/usr/bin/env bash

# export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ./data/test-1) #&& echo $WALLET_1;
# export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') #&& echo $ICA_ADDR
export WALLET_2=$(icad keys show wallet2 -a --keyring-backend test --home ./data/test-1) #&& echo $WALLET_2;
export ICA_ADDR_2=$(icad query intertx interchainaccounts connection-0 $WALLET_2 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') && echo $ICA_ADDR_2

export product="amazon.org"
export price="50"

echo ""
echo "Who is owner of the product $product ?"
echo "-----------------------------------------"


# icad tx controller submit-tx "{
#     \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
#     \"creator\": \"$ICA_ADDR_2\",
#     \"name\":\"$product\",
#     \"bid\":\"110\",
#     \"metadata\":\"Brought by ICA Wallet 2\"
# }" connection-0 --from $WALLET_2 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y -o json | jq -r '.txhash'


# icad q nameservice list-whois --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 -o json | jq '.whois[] | select(.name=="amazon.org")'

icad q nameservice list-whois -o json | jq '.whois[] | select(.name=="amazon.org")'
# icad q nameservice list-whois -o json | jq '.whois[] | select(.name=="$product")'

# icad q nameservice list-whois --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 -o json | jq '.whois[] | select(.name=="$product")'

#icad q nameservice list-whois --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 -o json | jq '.whois[] | select(.name=="testcontroller.com")'

echo ""

#python3 scripts/query_status.py -r "testcontroller.com:::$ICA_ADDR" -w $WALLET_1 -ica $ICA_ADDR