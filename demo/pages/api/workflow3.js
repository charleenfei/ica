export default function workflow3(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules

  const cmd = 'cd ~/cosmos/sbip/; echo "Product price range:" $(cat oracle/cmp_config.json |  jq \'.price_range.".org"[]\' ); \
  make workflow3_submit_transaction;';

  const output = execSync(cmd, { encoding: 'utf-8' });  // the default is 'buffer'
  const splitted = output.split(/\r?\n/);  
  const filtered = splitted.filter( e => {
    return e !== '';
  });

  res.status(200).json(output)
}