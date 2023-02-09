export default function getICA(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules

  const cmd = 'export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ~/cosmos/sbip/data/test-1); \
  export TX_HASH=$(icad tx intertx register --from $WALLET_1 --connection-id connection-0 --chain-id test-1 --home ~/cosmos/sbip/data/test-1 --node tcp://localhost:16657 --keyring-backend test --yes -o json | jq -r .txhash); \
  export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ~/cosmos/sbip/data/test-1 --node tcp://localhost:16657 -o json | jq -r .interchain_account_address); \
  echo "";\
  echo "Transaction Hash: " $TX_HASH; echo "";\
  echo "WALLET_1 Address:" $WALLET_1;  echo "";\
  echo "ICA Address:" $ICA_ADDR;  echo "";';

  const output = execSync(cmd, { encoding: 'utf-8' });  // the default is 'buffer'
  const splitted = output.split(/\r?\n/);  
  const filtered = splitted.filter( e => {
    return e !== '';
  });

  res.status(200).json(output)
}