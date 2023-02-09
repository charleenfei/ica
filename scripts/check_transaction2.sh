#!/usr/bin/env bash

export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ./data/test-1) #&& echo $WALLET_1;
export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') #&& echo $ICA_ADDR
# export domain="amit.com"
# export domain="block1.com"
# export domain="test.xyz"
export domain="google.xyz"
echo ""
echo "Who is owner of the product $domain ?"
echo "-----------------------------------------"


# icad q nameservice list-whois --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 -o json | jq '.whois[] | select(.name=="block.com")'

icad q nameservice list-whois --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 -o json | jq '.whois[]'

# icad q nameservice list-whois --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 -o json | jq '.whois[] | select(.name=="$domain")'

#icad q nameservice list-whois --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 -o json | jq '.whois[] | select(.name=="testcontroller.com")'

echo ""

#python3 scripts/query_status.py -r "testcontroller.com:::$ICA_ADDR" -w $WALLET_1 -ica $ICA_ADDR