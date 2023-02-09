export default function handler3(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules
  const cmd = 'icad status | jq ."SyncInfo"."latest_block_height" -r';
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