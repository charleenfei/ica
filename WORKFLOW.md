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
   + Create one interchain account (ICA) from controller chain (chain-id: `test-1`)
      ```sh
      icad tx intertx register --from $WALLET_1 --connection-id connection-0 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test --timeout-height 1000 -y
      ```
     Store the interchain account address in a variable with
     ```sh
      export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') && echo $ICA_ADDR
     ```
     Make sure the query shows an address. If it fails, restart the relayer and retry account creation. If it still fails, try removing `$HOME/.ica` folder and rebuilding

2. **Controller verification module**

   Only KYC-ed account can send crosschain tx through `controller` module. By default, all accounts are not verified, so, all crosschain transactions are rejected

   + In the terminal prepared, try sending a buy request
      ```sh
      icad tx controller submit-tx \
      "{
          \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
          \"creator\": \"$ICA_ADDR\",
          \"name\":\"testcontroller.com\",
          \"bid\":\"1500\",
          \"metadata\":\"test_meta_data\"
      }" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y
      ```
        + Check the relayer terminal, there is no crosschain transaction relayed, it is rejected at controller chain. 
        + Check the list item on host chain, It shoule be empty
          ```sh
          icad q nameservice list-whois #(optional parameters) --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657
          ```
   + Change `oracle/cmp_config.json` to set `kyc = true` for `WALLET_1`
   + Try the buy request again. Then show the oracle's decision with 
       ```sh
       python3 scripts/query_status.py -r "testcontroller.com:::$ICA_ADDR" -w $WALLET_1 -ica $ICA_ADDR
       ```

     This time it shows another message
       ```sh
       Query result:  REJECT::Bid 1500 is out of price range 100 -> 200 for domain .com
       ```
   + We can fix this by either changing the price bid in the transaction or price range in the offchain json file. Let's say you change the price from 1500 to 150
       ```sh
      icad tx controller submit-tx \
      "{
          \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
          \"creator\": \"$ICA_ADDR\",
          \"name\":\"testcontroller.com\",
          \"bid\":\"150\",
          \"metadata\":\"test_meta_data\"
      }" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y
       ```
       Now the transaction is successful!
       ```sh
        python3 scripts/query_status.py -r "testcontroller.com:::$ICA_ADDR" -w $WALLET_1 -ica $ICA_ADDR
        #Expected output: Query result:  OK::
       ```
       Check the list item on host chain with command `icad q nameservice list-whois`, it should show `testcontroller.com` belongs to interchain account of `$WALLET_1`


3. **Workflow A: Banned domain cannot be bought**

   In the default configuration, `.country-x` domain is banned
   + Try buying a `.country-x` domain with
       ```sh
      icad tx controller submit-tx \
      "{
          \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
          \"creator\": \"$ICA_ADDR\",
          \"name\":\"testdomain.country-x\",
          \"bid\":\"200\",
          \"metadata\":\"test_meta_data\"
      }" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y
       ```
       It will be rejected. Check with:
       ```sh
       icad q nameservice list-whois
       ## testdomain.country-x won't show up
       ```
   + To unban `.country-x` domain, remove `.country-x` from the banned entries in the offchain file `oracle/cmp_config.json`
   + Try buying the `.country-x` domain again, this time, it should be successful 
       ```sh
       icad q nameservice list-whois
       ## testdomain.country-x should belong to interchain account of $WALLET_1
       ```

4. **Workflow B: Price control, certain domain has certain price range**

   Example in `cmp_config.json`:  `".org": [10,20]` means domain `.org` can be bought with price between 10 and 20
   + With a similar command, try buying a `.org` domain with bid=50 (outside acceptable range):
       ```sh
      icad tx controller submit-tx \
      "{
          \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
          \"creator\": \"$ICA_ADDR\",
          \"name\":\"testdomain.org\",
          \"bid\":\"50\",
          \"metadata\":\"test_meta_data\"
      }" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y
       ```
       It will be rejected. Again, check with:
       ```sh
       icad q nameservice list-whois
       ## testdomain.org won't show up
       ## testcontroller.com and testdomain.country-x show up if you've done previous workflows
       ```

   + Change in `oracle/cmp_config.json`:  `".org": [10,200]`, meaning domain `.org` can be bought with price between 10 and 200
   + Submit the same transaction again, now it should be accepted. Check with
       ```sh
       icad q nameservice list-whois
       ## testdomain.org now belongs to interchain account of $WALLET_1
       ```
5. **Scenario O: Oracle is offline**

   If oracle is offline, transactions will be in a pending list and never be cleared
   + Close the oracle process `oracle/simple_oracle.py` earlier
   + Try to buy some domain
       ```sh
      icad tx controller submit-tx \
      "{
          \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
          \"creator\": \"$ICA_ADDR\",
          \"name\":\"my_domain.com\",
          \"bid\":\"150\",
          \"metadata\":\"test_meta_data\"
      }" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y
       ```
       Check the result:
       ```sh
       icad q nameservice list-whois
       ## It should show my_domain.com is not accepted
       ```
