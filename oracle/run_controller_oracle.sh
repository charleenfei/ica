export $(cat /src/app/data/oracle/docker.env | xargs)

python3 oracle/controller_oracle.py ${CMP_CHAIN_WS}
