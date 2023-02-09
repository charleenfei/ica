export default function workflow1(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules

  const cmd = 'cd ~/cosmos/sbip/; echo "KYC Status for Wallet 1:" $(cat oracle/cmp_config.json | jq .user_info.cosmos1m9l358xunhhwds0568za49mzhvuxx9uxre5tud.kyc); make workflow1_submit_transaction;';

  const output = execSync(cmd, { encoding: 'utf-8' });  // the default is 'buffer'
  const splitted = output.split(/\r?\n/);  
  const filtered = splitted.filter( e => {
    return e !== '';
  });

  res.status(200).json(output)
}