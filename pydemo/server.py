#!/usr/bin/env python3

import subprocess
import os
from pathlib import Path
import sys

from flask import Flask

app = Flask(__name__, static_folder='main', static_url_path='')
aliases = {}

@app.route('/secret')
def get_index():
    return 'Main page! \u0394 \N{GRINNING FACE} \N{WINKING FACE}'

@app.route('/prices')
def get_prices():
    with open('../oracle/prices.json') as f:
        s = f.read()
        print(s)
        return s

@app.route('/sells')
def get_sells():
    return subprocess.check_output(['bash', 'sells.sh'])

@app.route('/owns')
def get_owns():
    return subprocess.check_output(['bash', 'owns.sh'])

@app.route('/alias/<address>')
def get_alias(address):
    print(aliases)
    if address in aliases:
        return aliases[address]
    return address

@app.route('/register/<address>/<name>')
def get_register(address, name):
    try:
        icaBytes = subprocess.check_output(['bash', str(Path(os.getcwd()).parent.absolute()) + '/run/register', address])
#        icaBytes = subprocess.getoutput('bash ' + str(Path(os.getcwd()).parent.absolute()) + '/run/register ' + address)
    except subprocess.CalledProcessError as exc:
        return str(exc.output) + '!!'
    else:
        ica = icaBytes.decode(sys.stdout.encoding).strip()
        aliases[ica] = name
        #print (aliases)
        return ica
#    return subprocess.check_output(['bash', str(Path(os.getcwd()).parent.absolute()) + '/run/register', address])

if __name__ == "__main__":
    app.run(debug=True, port=5555)
