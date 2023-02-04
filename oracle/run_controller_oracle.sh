export $(cat docker.env | xargs)

python3 oracle/controller_oracle.py
