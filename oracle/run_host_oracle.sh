export $(cat docker.env | xargs)

python3 oracle/simple_oracle.py ws://chain-test-2:26657/websocket
