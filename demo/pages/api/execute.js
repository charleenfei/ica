export default function handler(req, res) {
  const execSync = require('child_process').execSync;
  // import { execSync } from 'child_process';  // replace ^ if using ES modules

  const { exec } = require("child_process");
  exec("ls -la", (error, stdout, stderr) => {
    if (error) {
        console.log(`error: ${error.message}`);
        return;
    }
    if (stderr) {
        console.log(`stderr: ${stderr}`);
        return;
    }
    console.log(`stdout: ${stdout}`);
});

  // const output = execSync('ls', { encoding: 'utf-8' });  // the default is 'buffer'
  // const splitted = output.split(/\r?\n/);  
  // const filtered = splitted.filter( e => {
  //   return e !== '';
  // });

  // res.status(200).json(JSON.stringify(filtered))
}