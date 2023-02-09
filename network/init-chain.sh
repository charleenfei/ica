#!/bin/bash

# set -x

BINARY=icad
CHAIN_DIR=./data

CHAINID=$1
ADDRESS=$2
P2PPORT=$3
RPCPORT=$4
RESTPORT=$5
ROSETTA=$6
VAL_MNEMONIC=$7
RLY_MNEMONIC=$8

echo "Removing previous data..."
rm -rf $CHAIN_DIR/$CHAINID &> /dev/null

# Add directories for the chain, exit if an error occurs
if ! mkdir -p $CHAIN_DIR/$CHAINID 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

echo "Initializing chain $CHAINID..."
$BINARY init test --home $CHAIN_DIR/$CHAINID --chain-id=$CHAINID

echo "Adding genesis accounts..."
echo "Adding new + $VAL_MNEMONIC +"
echo $VAL_MNEMONIC | $BINARY keys add val --home $CHAIN_DIR/$CHAINID --recover --keyring-backend=test
echo $WALLET_MNEMONIC_1 | $BINARY keys add wallet1 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend=test
echo $WALLET_MNEMONIC_2 | $BINARY keys add wallet2 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend=test
echo $WALLET_MNEMONIC_3 | $BINARY keys add wallet3 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend=test
echo $WALLET_MNEMONIC_4 | $BINARY keys add wallet4 --home $CHAIN_DIR/$CHAINID --recover --keyring-backend=test

echo $RLY_MNEMONIC | $BINARY keys add rly --home $CHAIN_DIR/$CHAINID --recover --keyring-backend=test

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID keys show val --keyring-backend test -a) 100000000000stake  --home $CHAIN_DIR/$CHAINID

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID keys show wallet1 --keyring-backend test -a) 100000000000stake  --home $CHAIN_DIR/$CHAINID
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID keys show wallet2 --keyring-backend test -a) 200000000000stake  --home $CHAIN_DIR/$CHAINID
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID keys show wallet3 --keyring-backend test -a) 300000000000stake  --home $CHAIN_DIR/$CHAINID
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID keys show wallet4 --keyring-backend test -a) 400000000000stake  --home $CHAIN_DIR/$CHAINID

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID keys show rly --keyring-backend test -a) 100000000000stake  --home $CHAIN_DIR/$CHAINID

echo "Creating and collecting gentx..."
$BINARY gentx val 7000000000stake --home $CHAIN_DIR/$CHAINID --chain-id $CHAINID --keyring-backend test
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID

echo "Changing defaults and ports in app.toml and config.toml files..."
sed -i -e 's#"tcp://0.0.0.0:26656"#"tcp://'"$ADDRESS"':'"$P2PPORT"'"#g' $CHAIN_DIR/$CHAINID/config/config.toml
sed -i -e 's#"tcp://127.0.0.1:26657"#"tcp://'"$ADDRESS"':'"$RPCPORT"'"#g' $CHAIN_DIR/$CHAINID/config/config.toml
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' $CHAIN_DIR/$CHAINID/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' $CHAIN_DIR/$CHAINID/config/config.toml
sed -i -e 's/index_all_keys = false/index_all_keys = true/g' $CHAIN_DIR/$CHAINID/config/config.toml
sed -i -e 's/enable = false/enable = true/g' $CHAIN_DIR/$CHAINID/config/app.toml
sed -i -e 's/swagger = false/swagger = true/g' $CHAIN_DIR/$CHAINID/config/app.toml
sed -i -e 's#"tcp://0.0.0.0:1317"#"tcp://'"$ADDRESS"':'"$RESTPORT"'"#g' $CHAIN_DIR/$CHAINID/config/app.toml
sed -i -e 's#":8080"#":'"$ROSETTA"'"#g' $CHAIN_DIR/$CHAINID/config/app.toml
