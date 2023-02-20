#!/bin/bash 
icad q nameservice list-whois -o json | jq -r '.whois'
