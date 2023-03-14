"""
Users of the controller chain querying the the status of the CMP request from the remote host chain
Submit tx and wait for the event emitted by the IBC module OnACK
"""

import io
import websocket
# from threading import Thread
import time
import sys
import json
import os
import pprint
from subprocess import check_output
from shlex import quote
import argparse
from contextlib import redirect_stdout

CONTROLLER_NODE = "tcp://localhost:16657"
REQUEST_ID = "NA"
OWNER_WALLET = "NA"
ICA_ADDRESS = "NA"
REQUEST_DATA_SEPARATOR = "::::"
CHAIN_ID= "NA"

# For websocket subscription of events
ws_params = {"jsonrpc": "2.0", "method": "subscribe", "id": 0, "params": {"query": "tm.event = 'Tx'"}}

# from module's keys.go, string constants of the cmp events and attributes
CmpResultPrefix = "cmp-result-request"
CmpResultId = CmpResultPrefix+".request-id"
CmpResultData = CmpResultPrefix+".request-data"

# when ws receive message, logic of parsing events is here
def on_message(ws, message):
    message = json.loads(message)
    # print(message)
    if "result" in message and "events" in message["result"]:
        cmp_event = {}
        match_event = False
#        print(message["result"]["events"])
#        print("----")
        for event, event_attribute in message["result"]["events"].items():
            if CmpResultPrefix in event:
                cmp_event[event] = event_attribute[0]
                # print(event, event_attribute[0])
            if CmpResultData in event:
                # print ("event data ", event_attribute[0])
                # parse the event data
                head, _ = event_attribute[0].split("::::::")
                request_id, request_data = head.split(REQUEST_DATA_SEPARATOR)
                if request_id == REQUEST_ID:
                    print("Query result: ", request_data)
                    match_event = True
        if match_event:
            ws.close()


# get tx command template for submitting the callback to the blockchain
def get_tx_command(request_id, chain_id, chain_home, owner_wallet):
    tx_data = {
        "@type":"/cosmos.interchainaccounts.nameservice.MsgQueryCmpStatus",
        "creator": f"{ICA_ADDRESS}",
        "request": f"{request_id}"
    }
    return (
        f"docker-compose run chain-test-1 icad tx controller submit-tx '{json.dumps(tx_data)}' "
        f"connection-0 --chain-id {chain_id} --home {chain_home} --keyring-backend test --from {owner_wallet} --node {CONTROLLER_NODE} -y"
    )

# utility to run arbitrary command
def run_sh(command):
    check_output(command, shell=True, universal_newlines=True, env=os.environ).strip()

# construct command and callback to the host_cmp module handler on blockchain
def execute_crosschain_query(request_id):
    tx_command = get_tx_command(
        request_id,
        CHAIN_ID,
        "/home/ubuntu/data/" + CHAIN_ID,
        OWNER_WALLET
    )
    # print(f"\n  Command: {tx_command}")
    f_temp = io.StringIO()
    with redirect_stdout(f_temp):
        run_sh(tx_command)

def on_error(ws, error):
    pass

def on_close(ws, close_status_code, close_msg):
    print("### closed ###")

def on_open(ws):
    ws.send(json.dumps(ws_params))
    execute_crosschain_query(REQUEST_ID)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-r", "--request", type=str, help="The id of the request to query", required=True)
    parser.add_argument("-w", "--owner-wallet", type=str, help="The owner wallet of the ICA", required=True)
    parser.add_argument("-ica", "--ica-address", type=str, help="The ICA address", required=True)
    parser.add_argument("-ws", "--web-socket", type=str,default="ws://localhost:16657/websocket", help="The host of the websocket")
    parser.add_argument("-n", "--node", type=str, default="tcp://localhost:16657", help="the blockchain node address for submitting tx")
    parser.add_argument("-c", "--chain", type=str, default="test-1", help="chain-id, e.g., test-1")
    args = parser.parse_args()
#    websocket.enableTrace(False)

    host = args.web_socket
    CONTROLLER_NODE = args.node

    CHAIN_ID = args.chain

    if CHAIN_ID == "test-1":
        host = "ws://localhost:16657/websocket"
        CONTROLLER_NODE = "tcp://chain-test-1:16657"

    if CHAIN_ID == "test-3":
        host = "ws://localhost:36657/websocket"
        CONTROLLER_NODE = "tcp://chain-test-3:36657"

    REQUEST_ID = args.request
    OWNER_WALLET = args.owner_wallet
    ICA_ADDRESS = args.ica_address

    ws = websocket.WebSocketApp(host, on_message=on_message, on_error=on_error, on_close=on_close)
    # disable verbose tracing
    ws.on_open = on_open
    # TODO: add timeout

    ws.run_forever()
