#!/usr/bin/env python3

import requests

from flask import Flask

app = Flask(__name__)

@app.route('/')
def get_index():
    return 'This is controller oracle!'

@app.route('/last-price/<product_name>')
def get_last_price(product_name):
    r = requests.get(url = f"http://localhost:5555/last-price/{product_name}")
    return r.text

if __name__ == "__main__":
    app.run(debug=True, port=7777)
