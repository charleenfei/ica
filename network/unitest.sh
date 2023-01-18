set -e

echo "************ INTERCHAIN ACCOUNT DEMO UNITEST ************"

export $(cat data/oracle/docker.env | xargs)

CONTROLLERCHAIN_ID=chain-test-1
HOSTCHAIN_ID=chain-test-2

echo "[EXECUTING] setup ICA account from controller chain (ChainId: test-1)"
txhash=$(echo y | icad tx intertx register --from $WALLET_1 --connection-id connection-0 --chain-id test-1 --home ./data/test-1 --node tcp://${CONTROLLERCHAIN_ID}:16657 --keyring-backend test --timeout-height 1000 --broadcast-mode block --output json | jq -r '.txhash' | sed 's/null//g' | xargs)
echo "[INFO] txhash: ${txhash}"