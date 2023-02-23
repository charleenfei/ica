set -e
echo "*********************************************************"
echo "************ INTERCHAIN ACCOUNT DEMO UNITEST ************"
echo "*********************************************************"

export $(cat data/oracle/docker.env | xargs)

CMP_CONTROLLER_CHAIN_ID=test-1
CMP_CONTROLLER_CHAIN_URL=tcp://chain-test-1:16657
CMP_NEW_CONTROLLER_CHAIN_ID=test-3
CMP_NEW_CONTROLLER_CHAIN_URL=tcp://chain-test-3:36657
CMP_HOST_CHAIN_ID=test-2
CMP_HOST_CHAIN_URL=tcp://chain-test-2:26657
CONNECTION_ID=connection-0

echo "[EXECUTING] setup ICA account from controller chain (ChainId: ${CMP_CONTROLLER_CHAIN_ID})..."
txhash=$(icad tx intertx register --from $WALLET_1 --connection-id ${CONNECTION_ID} --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test --timeout-height 1000 --broadcast-mode block -y --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Query the address of the interchain account..."
sleep 10
ICA_ADDR=$(icad query intertx interchainaccounts ${CONNECTION_ID} $WALLET_1 --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --output json | jq -r '.interchain_account_address')
echo "[INFO] interchain_account_address: ${ICA_ADDR}"

echo
echo "****** WORKFLOW Z: Controller verification module ******"
echo

echo "[EXECUTING] try sending a buy request WITHOUT KYC..."
sed -i '14s/true/false/' /oracle/cmp_config.json

txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"150\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 10

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
sleep 10

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

echo
echo "****** WORKFLOW A: Banned domain cannot be bought ******"
echo

echo "[INFO] In the default configuration, .country-x domain is banned..."
sed -i '2s/#/country-x/' /oracle/cmp_config.json
echo "[EXECUTING] Try buying a .country-x domain..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.country-x\",
    \"bid\":\"200\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')

echo "[INFO] txhash: ${txhash}"
sleep 10

expected_whois_size=$(icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.whois | length')

if [ ${expected_whois_size} -eq 1 ]; then
    echo
    echo "[SUCCESS!!!] DNS Item List is expected to have one record..."
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is NOT having one record..."
    icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json

    exit 1
fi

echo "[INFO] To unban .country-x domain"
sed -i '2s/country-x/#/' /oracle/cmp_config.json

echo "[EXECUTING] Re-try buying a .country-x domain..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.country-x\",
    \"bid\":\"200\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')

echo "[INFO] txhash: ${txhash}"
sleep 10

expected_whois_size=$(icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.whois | length')

if [ ${expected_whois_size} -eq 2 ]; then
    echo
    echo "[SUCCESS!!!] DNS Item List is expected to have two record..."
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is NOT having two record..."
    exit 1
fi

echo
echo "****** WORKFLOW B: Price control, certain domain has certain price range ******"
echo

echo "[EXECUTING] buying a testdomain.org domain with price = 50..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.org\",
    \"bid\":\"50\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')

echo "[INFO] txhash: ${txhash}"
sleep 10

expected_whois_size=$(icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.whois | length')
if [ ${expected_whois_size} -eq 2 ]; then
    echo
    echo "[SUCCESS!!!] DNS Item List is expected STILL to have two record..."
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is NOT having two record..."
    exit 1
fi

echo "[EXECUTING] buying a testdomain.org domain with price = 15..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testdomain.org\",
    \"bid\":\"15\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')

echo "[INFO] txhash: ${txhash}"
sleep 10

expected_whois_size=$(icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.whois | length')
if [ ${expected_whois_size} -eq 3 ]; then
    echo
    echo "[SUCCESS!!!] DNS Item List is expected to have three records..."
    echo
else
    echo
    echo "[ERROR!!!] DNS Item List is NOT having three records..."
    exit 1
fi

echo
echo "****** WORKFLOW C: Match buyer and seller ******"
echo

echo "[EXECUTING] Create an interchain account for $WALLET_2 on chain ${CMP_NEW_CONTROLLER_CHAIN_ID} and store its account address in variable $ICA_ADDR_2..."
txhash=$(icad tx intertx register --from $WALLET_2 --connection-id ${CONNECTION_ID} --chain-id ${CMP_NEW_CONTROLLER_CHAIN_ID} --home ./data/${CMP_NEW_CONTROLLER_CHAIN_ID} --node ${CMP_NEW_CONTROLLER_CHAIN_URL} --keyring-backend test --timeout-height 1000 --broadcast-mode block -y --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"

echo "[INFO] Query the address of the NEW interchain account..."
sleep 10
ICA_ADDR_2=$(icad query intertx interchainaccounts ${CONNECTION_ID} $WALLET_2 --home ./data/${CMP_NEW_CONTROLLER_CHAIN_ID} --node ${CMP_NEW_CONTROLLER_CHAIN_URL} --output json | jq -r '.interchain_account_address')
echo "[INFO] NEW interchain_account_address: ${ICA_ADDR_2}"

echo "[EXECUTING] Fund the interchain accounts"
txhash=$(icad tx bank send $WALLET_3 $ICA_ADDR 10000stake --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --keyring-backend test --timeout-height 1000 --broadcast-mode block -y --output json | jq -r '.txhash')

denom=$(icad q bank balances $ICA_ADDR --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.balances[].denom')
balance=$(icad q bank balances $ICA_ADDR --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.balances[].amount')
if [ ${denom} = "stake" ]; then
    echo
    echo "[SUCCESS!!!] \$ICA_ADDR is expected to have correct denom"
    echo
else
     echo
    echo "[ERROR!!!] Incorrect denom for \$ICA_ADDR!!!"
    exit 1
fi

if [ $balance -eq 10000 ]; then
    echo
    echo "[SUCCESS!!!] \$ICA_ADDR is expected to have correct balance"
    echo
else
     echo
    echo "[ERROR!!!] Incorrect balance for \$ICA_ADDR!!!"
    exit 1
fi

txhash=$(icad tx bank send $WALLET_3 $ICA_ADDR_2 10000stake --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --keyring-backend test --timeout-height 1000 --broadcast-mode block -y --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
denom=$(icad q bank balances $ICA_ADDR_2 --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.balances[].denom')
balance=$(icad q bank balances $ICA_ADDR_2 --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.balances[].amount')

if [ ${denom} = "stake" ]; then
    echo
    echo "[SUCCESS!!!] \$ICA_ADDR_2 is expected to have correct denom"
    echo
else
     echo
    echo "[ERROR!!!] Incorrect denom for \$ICA_ADDR_2!!!"
    exit 1
fi

if [ $balance -eq 10000 ]; then
    echo
    echo "[SUCCESS!!!] \$ICA_ADDR_2 is expected to have correct balance"
    echo
else
     echo
    echo "[ERROR!!!] Incorrect balance for \$ICA_ADDR_2!!!"
    exit 1
fi

pending_sell_size=$(icad q nameservice list-pending-sell --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.pendingSell | length')
if [ ${pending_sell_size} -eq 0 ]; then
    echo
    echo "[SUCCESS!!!] Pending sell list is expected to be empty..."
    echo
else
    echo
    echo "[ERROR!!!] Pending sell list is NOT empty..."
    exit 1
fi

echo "[EXECUTING] Put testcontroller.com up for sale..."
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpSell\",
    \"creator\": \"$ICA_ADDR\",
    \"name\":\"testcontroller.com\",
    \"price\":\"110\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_1 --chain-id ${CMP_CONTROLLER_CHAIN_ID} --home ./data/${CMP_CONTROLLER_CHAIN_ID} --node ${CMP_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 10

pending_sell_size=$(icad q nameservice list-pending-sell --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.pendingSell | length')
if [ ${pending_sell_size} -eq 1 ]; then
    echo
    echo "[SUCCESS!!!] Pending sell list is expected to have one item..."
    echo
else
    echo
    echo "[ERROR!!!] Pending sell list is NOT expected to be empty..."
    exit 1
fi

echo "[EXECUTING] \$WALLET_2 can buy testcontroller.com from \$WALLET_1 with 110 stake:"
txhash=$(icad tx controller submit-tx \
"{
    \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
    \"creator\": \"$ICA_ADDR_2\",
    \"name\":\"testcontroller.com\",
    \"bid\":\"110\",
    \"metadata\":\"test_meta_data\"
}" ${CONNECTION_ID} --from $WALLET_2 --chain-id ${CMP_NEW_CONTROLLER_CHAIN_ID} --home ./data/${CMP_NEW_CONTROLLER_CHAIN_ID} --node ${CMP_NEW_CONTROLLER_CHAIN_URL} --keyring-backend test -y --broadcast-mode block --output json | jq -r '.txhash')
echo "[INFO] txhash: ${txhash}"
sleep 10

pending_sell_size=$(icad q nameservice list-pending-sell --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.pendingSell | length')
if [ ${pending_sell_size} -eq 0 ]; then
    echo
    echo "[SUCCESS!!!] Pending sell list is expected to be empty..."
    echo
else
    echo
    echo "[ERROR!!!] Pending sell list is NOT empty..."
    exit 1
fi

pending_buy_size=$(icad q nameservice list-pending-buy --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.pendingSell | length')
if [ ${pending_buy_size} -eq 0 ]; then
    echo
    echo "[SUCCESS!!!] Pending buy list is expected to be empty..."
    echo
else
    echo
    echo "[ERROR!!!] Pending buy list is NOT empty..."
    exit 1
fi

owner=$(icad q nameservice list-whois --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.whois[] | select(.index == "testcontroller.com") | .owner')
if [ ${owner} = ${ICA_ADDR_2} ]; then
    echo
    echo "[SUCCESS!!!] \$ICA_ADDR_2 is new owner of testcontroller.com..."
    echo
else
    echo
    echo "[ERROR!!!] \$ICA_ADDR_2 is NOT new owner of testcontroller.com..."
    echo "owner: $owner"
    echo "ICA_ADDR_2: $ICA_ADDR_2"
    exit 1
fi

balance=$(icad q bank balances $ICA_ADDR --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.balances[].amount')
if [ $balance -eq 10110 ]; then
    echo
    echo "[SUCCESS!!!] \$ICA_ADDR is expected to have correct balance"
    echo
else
     echo
    echo "[ERROR!!!] Incorrect balance for \$ICA_ADDR!!!"
    exit 1
fi

balance=$(icad q bank balances $ICA_ADDR_2 --chain-id ${CMP_HOST_CHAIN_ID} --home ./data/${CMP_HOST_CHAIN_ID} --node ${CMP_HOST_CHAIN_URL} --output json | jq -r '.balances[].amount')
if [ $balance -eq 9890 ]; then
    echo
    echo "[SUCCESS!!!] \$ICA_ADDR_2 is expected to have correct balance"
    echo
else
     echo
    echo "[ERROR!!!] Incorrect balance for \$ICA_ADDR_2!!!"
    exit 1
fi
