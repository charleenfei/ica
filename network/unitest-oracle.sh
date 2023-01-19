set -e

export $(cat data/oracle/docker.env | xargs)

CONTROLLERCHAIN_ID=test-1
CONTROLLERCHAIN_URL=chain-test-1
HOSTCHAIN_ID=test-2
HOSTCHAIN_URL=chain-test-2
CONNECTION_ID=connection-0

echo
echo "****** WORKFLOW 3: If oracle is offline. Tx will be in pending list and never be cleared ******"
echo

ICA_ADDR=$(icad query intertx interchainaccounts ${CONNECTION_ID} $WALLET_1 --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --output json | jq -r '.interchain_account_address')
echo "[INFO] interchain_account_address: ${ICA_ADDR}"

echo "[EXECUTING] Submit ICA the to register a name domain \"my_domain.com\"..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"my_domain.com\",
    \"bid\":\"50\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Show that it's not accepted, only 3 domains: \"testcontroller.com\", \"testdomain.country-x\" were registered..."
icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json

echo "[INFO] Show that there is a pending item waiting for cmp result: item \"my_domain.com\"..."
icad q nameservice list-pending-buy --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json