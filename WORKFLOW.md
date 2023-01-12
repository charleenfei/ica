# Workflow

:warning: **WARNING:** SDK modules on a chain are assumed to be trustworthy. For example, there are no checks to prevent an untrustworthy module from accessing the bank keeper.
(Quoted from https://ibc.cosmos.network/main/apps/interchain-accounts/overview.html)

## Overview 

This demo follows this workflow

### Setup
1. **Build the `icad` binary, the relayer binary, and run an `icad` node and a relayer**

   Run the following and keep the relayer running in that terminal

    ```sh
    make install
    make init-golang-rly
    make start-golang-rly
    ```

2. **Run the oracle**

   + Open a new terminal

   + Install pip package if not yet installed

      ```sh
      pip3 install -r oracle/requirements.txt
      ```
   + Run the oracle

      ```sh
      make start-oracle
      ```
   + Keep it running

### Workflows
1. **Preparation**
   + Open a new terminal
   + Setup account addresses in the shell env

      ```sh
      export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ./data/test-1) && echo $WALLET_1;
      export WALLET_2=$(icad keys show wallet2 -a --keyring-backend test --home ./data/test-1) && echo $WALLET_2;
      export WALLET_3=$(icad keys show wallet3 -a --keyring-backend test --home ./data/test-2) && echo $WALLET_3;
      export WALLET_4=$(icad keys show wallet4 -a --keyring-backend test --home ./data/test-2) && echo $WALLET_4;
      ```
   + Create one ICA from controller chain (chain-id: `test-1`)
      ```sh
      icad tx intertx register --from $WALLET_1 --connection-id connection-0 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test --timeout-height 1000 -y
      ```
     Store the interchain account address in a variable with
     ```sh
      export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') && echo $ICA_ADDR
     ```
     Make sure the query shows an address. If it fails, restart the relayer and retry account creation. If it still fails, try removing `$HOME/.ica` folder and rebuilding

2. **Controller verification module**

   Only KYC-ed account can send crosschain tx through `controller` module
