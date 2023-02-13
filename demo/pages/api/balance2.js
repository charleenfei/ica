export default function balance(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules


  const cmd = 'echo ""; \
  export WALLET_2=$(icad keys show wallet2 -a --keyring-backend test --home ~/cosmos/sbip/data/test-1);\
  export ICA_ADDR_2=$(icad query intertx interchainaccounts connection-0 $WALLET_2 --home ~/cosmos/sbip/data/test-1 --node tcp://localhost:16657 -o json | jq -r .interchain_account_address);\
  echo "WALLET_2 Address:" $WALLET_2;  echo "";\
  echo "ICA Address 2:" $ICA_ADDR_2;  echo "";\
  export Balance2=$(icad q bank balances $ICA_ADDR_2 --chain-id test-2 --home ~/cosmos/sbip/data/test-2 -o json |  jq -r \'.balances[].amount\');\
  echo "Balance of ICA Address 2: $Balance2";\
  echo "";'

  // const cmd = 'echo ""; \
  // export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ~/cosmos/sbip/data/test-1); \
  // export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ~/cosmos/sbip/data/test-1 --node tcp://localhost:16657 -o json | jq -r .interchain_account_address); \
  // echo "WALLET_1 Address:" $WALLET_1;  echo "";\
  // echo "ICA Address 1:" $ICA_ADDR;  echo "";\
  // echo "Balance of ICA Address 1:";\
  // icad q bank balances $ICA_ADDR --chain-id test-2 --home ~/cosmos/sbip/data/test-2 -o json |  jq -r \'.balances[].amount\';\
  // echo "-----------------------------------------";\
  // echo "WALLET_2 Address:" $WALLET_2;  echo "";\
  // echo "ICA Address 2:" $ICA_ADDR_2;  echo "";\
  // echo "Balance of ICA Address 2:";\
  // icad q bank balances $ICA_ADDR_2 --chain-id test-2 --home ~/cosmos/sbip/data/test-2 -o json |  jq -r \'.balances[].amount\';\
  // echo "-----------------------------------------";\
  // echo "";'
  const output = execSync(cmd, { encoding: 'utf-8' });  // the default is 'buffer'
  const splitted = output.split(/\r?\n/);  
  const filtered = splitted.filter( e => {
    return e !== '';
  });

  // res.status(200).json(JSON.stringify(filtered))
  // res.status(200).json(filtered)
  // res.status(200).json(splitted)
  res.status(200).json(output)
}