export $(cat /src/app/data/oracle/docker.env | xargs)

python3 oracle/controller_oracle.py ${CMP_CHAIN_WS} &
CMP_CONTROLLER_CHAIN_ID=test-3 CMP_CHAIN_WS=ws://chain-test-3:36657/websocket CMP_CONTROLLER_CHAIN_HOME=/src/app/data/test-3 python3 oracle/controller_oracle.py ws://chain-test-3:36657/websocket
