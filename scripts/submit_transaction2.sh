#!/usr/bin/env bash

export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ./data/test-1) #&& echo $WALLET_1;
export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') #&& echo $ICA_ADDR

# export domain="testcontroller1.com"
export domain="test.xyz"
echo ""
echo "Buying the product name $domain "
echo "-----------------------------------------"

# icad tx controller submit-tx \
# "{
#     \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
#     \"creator\": \"$ICA_ADDR\",
#     \"name\":\"$domain\",
#     \"bid\":\"150\",
#     \"metadata\":\"test_meta_data\"
# }" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y -o json | jq -r '.txhash'

icad tx controller submit-tx \
"{
   \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
   \"creator\": \"$ICA_ADDR\",
   \"name\":\"$domain\",
   \"bid\":\"200\",
   \"metadata\":\"test_meta_data\"
}" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y -o json | jq -r '.txhash'