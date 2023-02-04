set -e
echo "*********************************************************"
echo "************ INTERCHAIN ACCOUNT DEMO UNITEST ************"
echo "*********************************************************"

export $(cat data/oracle/docker.env | xargs)

CMP_CONTROLLER_CHAIN_ID=test-1
CMP_CONTROLLER_CHAIN_URL=tcp://chain-test-1:16657
CMP_HOST_CHAIN_ID=test-2
CMP_HOST_CHAIN_URL=tcp://chain-test-2:26657
CONNECTION_ID=connection-0

echo "[EXECUTING] setup ICA account from controller chain (ChainId: ${CMP_CONTROLLER_CHAIN_ID})..."
txhash=$(icad tx intertx register --from $WALLET_1 --connection-id ${CONNECTION_ID} --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test --timeout-height 1000 --broadcast-mode block -y --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Query the address of the interchain account..."
sleep 5
ICA_ADDR=$(icad query intertx interchainaccounts ${CONNECTION_ID} $WALLET_1 --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --output json | jq -r '.interchain_account_address')
echo "[INFO] interchain_account_address: ${ICA_ADDR}"

echo
echo "****** WORKFLOW Z: Controller verification module ******"
echo

echo "[EXECUTING] try sending a buy request WITHOUT KYC..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"150\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

expected_whois_size=$(icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.whois | length')

if [ ${expected_whois_size} -eq 0 ]; then
    echo
    echo "[SUCCESS!!!] DNS Item List is expected to be be empty..."
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is NOT empty..."
    exit 1
fi

echo "[INFO] Change oracle/cmp_config.json to set kyc = true for WALLET_1"
sed -i '14s/false/true/' /oracle/cmp_config.json

echo "[EXECUTING] try RESENDING the buy request WITH KYC enabled now..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"150\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

expected_whois_size=$(icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.whois | length')

if [ ${expected_whois_size} -eq 1 ]; then
    echo
    echo "[SUCCESS!!!] DNS Item List is expected NOT to be be empty..."
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is STILL empty..."
    exit 1
fi