export $(cat /src/app/data/oracle/docker.env | xargs)

python3 oracle/host_oracle.py ${CMP_CHAIN_WS}
