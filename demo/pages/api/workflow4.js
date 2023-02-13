export default function workflow4(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules

  const cmd = 'cd ~/cosmos/sbip/; \
  echo "Wallet Balances";\
  make query_balance addr="cosmos19a50wg5l6tzyq9w9sxezqctxaud5v9ak2e0gnlp2ukl20x2j773qww6atd" chain=test-2;\
  make query_balance addr="cosmos1rm99r4z7njy584y80r0a584dawpaw86hw7qqwq9jkp6sx439f7pssp7qcs" chain=test-2;\
  make workflow4_submit_transaction; \
  echo ""; ';

  // icad q nameservice list-pending-sell -o json | jq \'.pendingSell[]\' ; \

  const output = execSync(cmd, { encoding: 'utf-8' });  // the default is 'buffer'
  const splitted = output.split(/\r?\n/);  
  const filtered = splitted.filter( e => {
    return e !== '';
  });

  res.status(200).json(output)
}