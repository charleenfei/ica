export $(cat docker.env | xargs)

python3 oracle/host_oracle.py
