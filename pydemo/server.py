#!/usr/bin/env python3

import subprocess

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
    return address

if __name__ == "__main__":
    app.run(debug=True, port=5555)
