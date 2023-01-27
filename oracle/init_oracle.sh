# Store the following account addresses within the current shell env
export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ./data/test-1) && echo $WALLET_1;
export WALLET_2=$(icad keys show wallet2 -a --keyring-backend test --home ./data/test-1) && echo $WALLET_2;
export WALLET_3=$(icad keys show wallet3 -a --keyring-backend test --home ./data/test-2) && echo $WALLET_3;
export WALLET_4=$(icad keys show wallet4 -a --keyring-backend test --home ./data/test-2) && echo $WALLET_4;

export CMP_ORACLE_WALLET=$(icad keys show wallet4 -a --keyring-backend test --home ./data/test-2) && echo $CMP_ORACLE_WALLET;

mkdir -p /home/ubuntu/data/oracle

echo "WALLET_1=${WALLET_1}" > /home/ubuntu/data/oracle/docker.env
echo "WALLET_2=${WALLET_2}" >> /home/ubuntu/data/oracle/docker.env
echo "WALLET_3=${WALLET_3}" >> /home/ubuntu/data/oracle/docker.env
echo "WALLET_4=${WALLET_4}" >> /home/ubuntu/data/oracle/docker.env
echo "CMP_ORACLE_WALLET=${CMP_ORACLE_WALLET}" >> /home/ubuntu/data/oracle/docker.env