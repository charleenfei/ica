set -e
echo "*********************************************************"
echo "************ INTERCHAIN ACCOUNT DEMO UNITEST ************"
echo "*********************************************************"

export $(cat data/oracle/docker.env | xargs)

CONTROLLERCHAIN_ID=test-1
CONTROLLERCHAIN_URL=chain-test-1
HOSTCHAIN_ID=test-2
HOSTCHAIN_URL=chain-test-2
CONNECTION_ID=connection-0

echo "[EXECUTING] setup ICA account from controller chain (ChainId: ${CONTROLLERCHAIN_ID})..."
txhash=$(icad tx intertx register --from $WALLET_1 --connection-id ${CONNECTION_ID} --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test --timeout-height 1000 --broadcast-mode block -y --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Query the address of the interchain account..."
sleep 5
ICA_ADDR=$(icad query intertx interchainaccounts ${CONNECTION_ID} $WALLET_1 --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --output json | jq -r '.interchain_account_address')
echo "[INFO] interchain_account_address: ${ICA_ADDR}"

echo "[EXECUTING] setup ICA account from controller chain (ChainId: ${CONTROLLERCHAIN_ID})..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"150\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

expected_whois_size=$(icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json | jq -r '.whois | length')

if [ ${expected_whois_size} -eq 0 ]; then
    echo
    echo "[SUCCESS!!!] DNS Item List is expected to be be empty..."
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is NOT empty..."
    exit 1
fi

expected_cmpData_size=$(icad q controller list-cmp-data --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --output json | jq -r '.cmpData | length')
if [ ${expected_cmpData_size} -eq 0 ]; then
    echo
    echo "[SUCCESS!!!] cmp data on controller chain is expected to be blank"
    echo
else
    echo
    echo "[ERROR!!!] cmp data on controller chain is NOT empty..."
    exit 1
fi

echo "[EXECUTING] Simulate offchain process submits KYC info from offchain source for WALLET_1..."
txhash=$(icad tx controller cmp-controller-push $WALLET_1 true retail test_meta_data --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

expected_cmpData_size=$(icad q controller list-cmp-data --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --output json | jq -r '.cmpData | length')
if [ ${expected_cmpData_size} -eq 1 ]; then
    echo
    echo "[SUCCESS!!!] After submit KYC, cmp data on controller chain is NOT empty"
    echo
else
    echo
    echo "[ERROR!!!] cmp data on controller chain is empty..."
    exit 1
fi

echo "[EXECUTING] Re-submit the same Tx again to setup ICA account from controller chain (ChainId: ${CONTROLLERCHAIN_ID})..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"150\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

echo "[INFO] Check the List item on Hostchain, verify there is only one name registered \"testcontroller.com\" (from the controller chain test)..."
expected_whois_size=$(icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json | jq -r '.whois | length')

if [ ${expected_whois_size} -eq 1 ]; then
    echo
    echo "[SUCCESS!!!] testcontroller.com successfully registered"
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is STILL empty..."
    exit 1
fi

echo
echo "****** WORKFLOW 1: Banned Domain Cannot Be Bought ******"
echo

echo "[EXECUTING] Submit ICA tx to register a name domain \".country-x\" under ICA address with bid = 200..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.country-x\",
    \"bid\":\"200\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

echo "[INFO] show that it's not showed on list-whois..."
icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json

echo "[EXECUTING] Re-submit ICA the same tx to register a name domain \".country-x\" again, this time it is not banned..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.country-x\",
    \"bid\":\"200\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

echo "[INFO] Show that the name is bought..."
icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json

echo
echo "****** WORKFLOW 2: Price control, certain domain has certain price range ******"
echo

echo "[EXECUTING] Submit ICA the same tx to register a name domain \".org\" with bid = 50, out side acceptable range..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.org\",
    \"bid\":\"50\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

echo "[INFO] Show that it's not accepted, only 2 domains: \"testcontroller.com\", \"testdomain.country-x\" were registered..."
icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json

echo "[EXECUTING] submit the same tx again, this time it should be accepted..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.org\",
    \"bid\":\"50\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CONTROLLERCHAIN_ID} --home ./data/${CONTROLLERCHAIN_ID} --node tcp://${CONTROLLERCHAIN_URL}:16657 --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 5

echo "[INFO] show that it's accepted, 3 domains \"testcontroller.com\", \"testdomain.country-x\" and \"testdomain.org\" were registered..."
icad q nameservice list-whois --chain-id ${HOSTCHAIN_ID} --home ./data/${HOSTCHAIN_ID} --node tcp://${HOSTCHAIN_URL}:26657 --output json