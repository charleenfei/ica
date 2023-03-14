#!/usr/bin/env python3

import subprocess
import os
from pathlib import Path
import sys
import json

from flask import Flask

app = Flask(__name__, static_folder='main', static_url_path='')
aliases = {}

source = os.path.dirname(__file__)
parent = os.path.join(source, '../')

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

@app.route('/balance/<address>')
def get_balance(address):
    balanceStr = subprocess.check_output(['bash', 'balance.sh', address])
    balanceObj = json.loads(balanceStr)
    res = ""
    for balance in balanceObj["balances"]:
        res += str(balance["amount"]) + " " + balance["denom"] + ","

    if res == "":
        res = "Empty!"
    return res

@app.route('/register/<address>/<name>/<chain>')
def get_register(address, name, chain):
    try:
        icaBytes = subprocess.check_output(['bash', str(Path(os.getcwd()).parent.absolute()) + '/run/register', address, chain])
#        icaBytes = subprocess.getoutput('bash ' + str(Path(os.getcwd()).parent.absolute()) + '/run/register ' + address + ' ' + chain)
    except subprocess.CalledProcessError as exc:
        return str(exc.output) + '!!'
    else:
        ica = icaBytes.decode(sys.stdout.encoding).strip()
        aliases[ica] = name
        #print (aliases)
        return ica
#    return subprocess.check_output(['bash', str(Path(os.getcwd()).parent.absolute()) + '/run/register', address])

@app.route('/run/<com>/<addr>/<chain>/<item>/<price>')
def get_run(com, addr, chain, item, price):
    command = 'bash ' + parent +'/run/' + com + ' ' + addr + ' ' + chain + ' ' + item + ' ' + price
    print(command)
    s = subprocess.getoutput(command)
    print(s)
    return s

if __name__ == "__main__":
    app.run(host='0.0.0.0', debug=True, port=5555)
