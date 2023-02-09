export default function handler2(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules
  const cmd = 'icad q ibc channel channels --home ../data/test-1 --node tcp://localhost:16657';
  const output = execSync(cmd, { encoding: 'utf-8' });  // the default is 'buffer'
  // const splitted = output.split(/\r?\n/);  
  // const filtered = splitted.filter( e => {
  //   return e !== '';
  // });

  // res.status(200).json(JSON.stringify(filtered))
  // res.status(200).json(filtered)
  // res.status(200).json(splitted)
  res.json(output.toString())
}