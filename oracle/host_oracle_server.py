#!/usr/bin/env python3

import os
import json

from flask import Flask

app = Flask(__name__)

HOST_DATA_FILE = "oracle/host-data.json"

@app.route('/')
def get_index():
    return 'Main page! \u0394 \N{GRINNING FACE} \N{WINKING FACE}'

@app.route('/last-price/<product_name>')
def get_last_price(product_name):
    data_path = os.getcwd() + "/" + HOST_DATA_FILE
    data = json.load(open(data_path, "r"))
    if not "last_price" in data:
        data["last_price"] = {}

    if not product_name in data["last_price"]:
        return "Unknown"

    return data["last_price"][product_name]

if __name__ == "__main__":
    app.run(debug=True, port=5555)
