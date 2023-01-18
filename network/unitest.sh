export $(cat data/oracle/docker.env | xargs)

echo "setup ICA account from controller chain (ChainId: test-1)"
echo y | icad tx intertx register --from $WALLET_1 --connection-id connection-0 --chain-id test-1 --home ./data/test-1 --node tcp://chain-test-1:16657 --keyring-backend test --timeout-height 1000