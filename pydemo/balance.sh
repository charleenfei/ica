#!/bin/bash

# set -x

ADDRESS=$1
icad q bank balances $ADDRESS -o json
