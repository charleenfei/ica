# Workflow

:warning: **WARNING:** SDK modules on a chain are assumed to be trustworthy. For example, there are no checks to prevent an untrustworthy module from accessing the bank keeper.
(Quoted from https://ibc.cosmos.network/main/apps/interchain-accounts/overview.html)

## Overview 

This demo follows this workflow

### Setup

- Reset data: `make docker-reset`

- Build icad, relayer & oracle docker images: `make docker-build`

- Init blockchain data : `make docker-init-chain`

- Start blockchain: `make docker-start-chain`

- Init relayer data: `make docker-init-relayer`

- Start relayer: `make docker-start-relayer`

- Init oracle wallet: `make docker-init-oracle`

- Start oracle: `make docker-start-oracle`

- Open `storage/oracle/docker.env`, verify the following account addresses
```
WALLET_1
WALLET_2
WALLET_3
WALLET_4
CMP_ORACLE_WALLET
```

- Run unitest: `make docker-unitest`

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

2. **Workflow Z: Controller verification module**

   Only KYC-ed account can send crosschain tx through `controller` module. By default, all accounts are not verified, so, all crosschain transactions are rejected

   + In the terminal prepared, try sending a buy request
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
        + Check the relayer terminal, there is no crosschain transaction relayed, it is rejected at controller chain. 
        + Check the list item on host chain, It shoule be empty
          ```sh
          icad q nameservice list-whois #(optional parameters) --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657
          ```
   + Change `oracle/cmp_config.json` to set `kyc = true` for `WALLET_1`
       ```sh
       sed -i '14s/false/true/' oracle/cmp_config.json  # change false -> true on line 14 of the file
       ```
   + Try the buy request again. Now the transaction is successful!
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

   Example in `oracle/cmp_config.json`:  `".org": [10,20]` means domain `.org` can be bought with price between 10 and 20
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

   + Submit the same transaction, but with bid 15, it should be accepted. Check with
       ```sh
       icad q nameservice list-whois
       ## testdomain.org now belongs to interchain account of $WALLET_1
       ```

5. **Workflow C: Match buyer and seller**

   Assume you have done Workflow Z, now `testcontroller.com` belongs to interchain account of `$WALLET_1`

   In this workflow, `$WALLET_2` will buy `testcontroller.com` from `$WALLET_1`.

   + Create an interchain account for `$WALLET_2` and store its account address in variable `$ICA_ADDR_2`. Make sure this step is successful, try restarting relayer if it is not
       ```sh
       icad tx intertx register --from $WALLET_2 --connection-id connection-0 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y
       export ICA_ADDR_2=$(icad query intertx interchainaccounts connection-0 $WALLET_2 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') && echo $ICA_ADDR_2
       ```

   + Fund the interchain accounts
       ```sh
       icad tx bank send $WALLET_3 $ICA_ADDR 10000stake --chain-id test-2 --home ./data/test-2 --keyring-backend test -y
       icad tx bank send $WALLET_3 $ICA_ADDR_2 10000stake --chain-id test-2 --home ./data/test-2 --keyring-backend test -y
       ```
       Now `ICA_ADDR` and `ICA_ADDR_2` each should have 10,000 stakes. Check with
       ```sh
       icad q bank balances $ICA_ADDR
       icad q bank balances $ICA_ADDR_2
       ```
   + `$WALLET_1` puts the name up for sale.

     First, check that there is no item for sale with `icad q nameservice list-pending-sell`

     Next, `$WALLET_1` submits a transaction to put `testcontroller.com` up for sale with:
       ```sh
       icad tx controller submit-tx "{
          \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpSell\",
          \"creator\": \"$ICA_ADDR\",
          \"name\":\"testcontroller.com\",
          \"price\":\"110\",
          \"metadata\":\"test_meta_data\"
       }" connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y
       ```
     Then check again list of items for sale with `icad q nameservice list-pending-sell`. The name should show up for sale.

   + Make sure `$WALLET_2` KYC is true in `oracle/cmp_config.json`. Then `$WALLET_2` can buy `testcontroller.com` from `$WALLET_1` with 110 stake:
     ```sh
     icad tx controller submit-tx "{
         \"@type\":\"/cosmos.interchainaccounts.nameservice.MsgCmpBuy\",
         \"creator\": \"$ICA_ADDR_2\",
         \"name\":\"testcontroller.com\",
         \"bid\":\"110\",
         \"metadata\":\"test\"
     }" connection-0 --from $WALLET_2 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y

     ```
   + Check that the transaction is successful
     ```sh
     icad q nameservice list-pending-sell   # should be empty
     icad q nameservice list-pending-buy    # should be empty
     icad q nameservice list-whois          # testcontroller.com should belong to $ICA_ADDR_2 
     #($WALLET_2 bought from $WALLET_1)
     icad q bank balances $ICA_ADDR         # ICA of $WALLET_1 should have 10110 stakes
     icad q bank balances $ICA_ADDR_2       # ICA of $WALLET_2 should have 9890 stakes
     ```

   + There are multiple scenarios that you can test with this workflow:
      * If the name is not owned, buyer gets it for free (TODO: Maybe change to default price of 100)
      * If the name is already owned but the owner is not selling, then buyer cannot buy
      * If the buy price is different from sell price, the transaction won't be successful
      * A seller cannot sell a name they don't own
      * A seller can sell at any price but a buyer has to buy at the price that the oracle approves (though they must match)
      * TODO: Let the seller stop selling the item or change price
