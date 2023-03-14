#!/bin/bash
icad q nameservice list-cmp-host-result -o json | jq -r '.cmpHostResult'
