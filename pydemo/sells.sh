#!/bin/bash 
icad q nameservice list-pending-sell -o json | jq -r '.pendingSell'
