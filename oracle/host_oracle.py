"""This is a simple implementation of oracle simulating Chainlink node's call back model
The reference request / response model:
https://docs.chain.link/architecture-overview/architecture-request-model

This oracle service will listen on cosmos blockchain websocket for events with CmpHostEventPrefix
When there is interesting events (CMP related), the oracle service will:
    1. Parse these event and read an off-chain json file in CMP_CONFIG_FILE
    2. Call check_cmp_logic() to check the request info against the off-chain cmp config -> OK/REJECT:
        - Proof of concept cmp logic : Banned list + price range for item categories (domain names)
        - Change of the CMP_CONFIG_FILE will take effect for all subsequent transactions
    3. The return YES/NO reply is submitted to the blockchain with host_cmp_callback()
"""
import websocket
# from threading import Thread
import time
import sys
import json
import os
import pprint
from subprocess import check_output
from shlex import quote

# file path for configuring CMP logic
CMP_CONFIG_FILE = "oracle/cmp_config.json"

# For websocket subscription of events
ws_params = {"jsonrpc": "2.0", "method": "subscribe", "id": 0, "params": {"query": "tm.event = 'Tx'"}}

# from module's keys.go, string constants of the cmp events and attributes
CmpHostEventPrefix = "cmp-host-request"
CmpHostCreator = "cmp-host-request.request-creator"
CmpHostId = "cmp-host-request.request-id"
CmpHostItem = "cmp-host-request.request-item"
CmpHostBid = "cmp-host-request.request-bid"
CmpHostMetaData = "cmp-host-request.request-metadata"

# when ws receive message, logic of parsing events is here
def on_message(ws, message):
    # print(message)
    print("\n\n" + "*" * 80)
    message = json.loads(message)
    if "result" in message and "events" in message["result"]:
        cmp_event = {}
        for event, event_attribute in message["result"]["events"].items():
            if CmpHostEventPrefix in event:
                cmp_event[event] = event_attribute[0]
                print(event, event_attribute[0])
            if "tx.hash" in event:
                print("Host Chain: ",event, event_attribute[0])
        if CmpHostId in cmp_event:

            # cmp event exist, process logic
            execute_cmp_logic(cmp_event)
    print("*" * 80 + "\n\n")


# when there is CMP event, execute the cmp logic on that event
def execute_cmp_logic(cmp_event):
    config_path = os.getcwd() + "/" + CMP_CONFIG_FILE
    try:
        config = json.load(open(config_path, "r"))
        print(f"Loaded config")
        pprint.pprint(config, indent=2)
        result, reason = check_cmp_logic(cmp_event, config)
        if result:
            host_cmp_callback(cmp_event[CmpHostId], "OK::" + reason)
            return
        else:
            host_cmp_callback(cmp_event[CmpHostId], "REJECT::" + reason)
            return
    except Exception as err:
        print("Failed to load config and execute CMP logic: ", err)
        print("Send NO to host-cmp-module")
        host_cmp_callback(cmp_event[CmpHostId], "REJECT::Exception when loading config and executing CMP logic")
        return


# check the cmp event against the config, return True/False + reason
def check_cmp_logic(cmp_event, cmp_config):
    domain_name = "." + cmp_event[CmpHostItem].split(".")[-1]
    bid = int(cmp_event[CmpHostBid])
    print(f"Checking Domain {domain_name}, bid {bid}")
    # check banned / sanction
    if "banned" in cmp_config:
        if "." + domain_name.split(".")[-1] in cmp_config["banned"]:
            reason = f"Domain {domain_name} is banned"
            return False, reason

    # check price range
    price_range = cmp_config["price_range"]["default"]
    if domain_name in cmp_config["price_range"]:
        price_range = cmp_config["price_range"][domain_name]

    if bid < price_range[0] or bid > price_range[1]:
        reason = f"Bid {bid} is out of price range {price_range[0]} -> {price_range[1]} for domain {domain_name}"
        return False, reason

    # Optional: extra logic with metadata
    # meta_data = cmp_event[CmpHostMetaData]

    return True, ""

# get tx command template for submitting the callback to the blockchain
def get_tx_command(request_id, decision, chain_id, chain_home, oracle_wallet):
    # print(" build tx command ", request_id, decision, chain_id, chain_home, oracle_wallet)
    return (
        f"icad tx nameservice cmp-host-callback {request_id} '{decision}' "
        f"--chain-id {chain_id} --home {chain_home} --keyring-backend test --from {oracle_wallet} -y"
    )

# utility to run arbitrary command
def run_sh(command):
    return check_output(command, shell=True, universal_newlines=True, env=os.environ).strip()

# construct command and callback to the host_cmp module handler on blockchain
def host_cmp_callback(request_id, decision):
    tx_command = get_tx_command(
        request_id,
        decision,
        os.environ.get("CMP_HOST_CHAIN_ID") or "test-2",
        os.environ.get("CMP_HOST_CHAIN_HOME") or os.getcwd() + "/data/test-2",
        os.environ.get("CMP_ORACLE_WALLET") or os.environ.get("WALLET_1"),
    )
    print(f"Host cmp callback:")
    print(f"\n  Request_id {request_id}")
    print(f"\n  Decision {decision}")
    # print(f"\n  Command: {tx_command}")
    run_sh(tx_command)


def on_error(ws, error):
    print(error)


def on_close(ws, close_status_code, close_msg):
    print("### closed ###")


def on_open(ws):
    ws.send(json.dumps(ws_params))

# Clear pending requests from blockchain
def clear_pending_buy():
    ...

if __name__ == "__main__":
    websocket.enableTrace(True)
    if len(sys.argv) < 2:
        host = "ws://localhost:26657/websocket"
    else:
        host = sys.argv[1]
    ws = websocket.WebSocketApp(host, on_message=on_message, on_error=on_error, on_close=on_close)
    # disable verbose tracing
    websocket.enableTrace(False)
    ws.on_open = on_open

    ws.run_forever()
