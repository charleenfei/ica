set -e

echo "************ INTERCHAIN ACCOUNT DEMO UNITEST ************"

export $(cat data/oracle/docker.env | xargs)

CONTROLLERCHAIN_ID=test-1
CONTROLLERCHAIN_URL=chain-test-1
HOSTCHAIN_ID=test-2
HOSTCHAIN_URL=chain-test-2

echo "[EXECUTING] setup ICA account from controller chain (ChainId: ${CONTROLLERCHAIN_ID})..."
txhash=$(icad tx intertx register --from $WALLET_1 --connection-id connection-0 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test --timeout-height 1000 --broadcast-mode block -y --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Query the address of the interchain account..."
sleep 2
ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --output json | jq -r '.interchain_account_address')
echo "[INFO] interchain_account_address: ${ICA_ADDR}"

echo "[EXECUTING] setup ICA account from controller chain (ChainId: ${CONTROLLERCHAIN_ID})..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"150\",
    \"metadata\":\"test_meta_data\"
}" connection-0 --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Check the List item on Hostchain, It shoule be empty..."
icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json

echo "[INFO] cmp data on controller chain is still bank, reject all crosschain tx..."
icad q controller list-cmp-data --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --output json

echo "[EXECUTING] Simulate offchain process submits KYC info from offchain source for WALLET_1..."
txhash=$(icad tx controller cmp-controller-push $WALLET_1 true retail test_meta_data --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

icad q controller list-cmp-data --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --output json

echo "[EXECUTING] Re-submit the same Tx again to setup ICA account from controller chain (ChainId: ${CONTROLLERCHAIN_ID})..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"150\",
    \"metadata\":\"test_meta_data\"
}" connection-0 --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Check the List item on Hostchain, verify there is only one name registered \"testcontroller.com\" (from the controller chain test)..."
icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json