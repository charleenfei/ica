"""
Users of the controller chain querying the the status of the CMP request from the remote host chain
Submit tx and wait for the event emitted by the IBC module OnACK
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
import argparse

CONTROLLER_NODE = "tcp://localhost:16657"
REQUEST_ID = "NA"
OWNER_WALLET = "NA"
ICA_ADDRESS = "NA"
REQUEST_DATA_SEPARATOR = "::::"
# For websocket subscription of events
ws_params = {"jsonrpc": "2.0", "method": "subscribe", "id": 0, "params": {"query": "tm.event = 'Tx'"}}

# from module's keys.go, string constants of the cmp events and attributes
CmpResultPrefix = "cmp-result-request"
CmpResultId = CmpResultPrefix+".request-id"
CmpResultData = CmpResultPrefix+".request-data"

# when ws receive message, logic of parsing events is here
def on_message(ws, message):
    # print(message)
    message = json.loads(message)
    if "result" in message and "events" in message["result"]:
        cmp_event = {}
        match_event = False
        for event, event_attribute in message["result"]["events"].items():
            if CmpResultPrefix in event:
                cmp_event[event] = event_attribute[0]
                # print(event, event_attribute[0])
            if CmpResultData in event:
                # print ("event data ", event_attribute[0])
                # parse the event data
                request_id, request_data = event_attribute[0].split(REQUEST_DATA_SEPARATOR)
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
        f"icad tx controller submit-tx '{json.dumps(tx_data)}' "
        f"connection-0 --chain-id {chain_id} --home {chain_home} --keyring-backend test --from {OWNER_WALLET} --node {CONTROLLER_NODE} -y"
    )

# utility to run arbitrary command
def run_sh(command):
    return check_output(command, shell=True, universal_newlines=True, env=os.environ).strip()

# construct command and callback to the host_cmp module handler on blockchain
def execute_crosschain_query(request_id):
    tx_command = get_tx_command(
        request_id,
        os.environ.get("CMP_CONTROLLER_CHAIN_ID") or "test-1",
        os.environ.get("CMP_CONTROLLER_CHAIN_HOME") or os.getcwd() + "/data/test-1",
        OWNER_WALLET
    )
    # print(f"\n  Command: {tx_command}")
    run_sh(tx_command)

def on_error(ws, error):
    print(error)

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
    args = parser.parse_args()
    websocket.enableTrace(False)
    host = args.web_socket
    REQUEST_ID = args.request
    CONTROLLER_NODE = args.node
    OWNER_WALLET = args.owner_wallet
    ICA_ADDRESS = args.ica_address
    ws = websocket.WebSocketApp(host, on_message=on_message, on_error=on_error, on_close=on_close)
    # disable verbose tracing
    ws.on_open = on_open
    # TODO: add timeout

    ws.run_forever()
