sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 8AA7AF1F1091A5FD
sudo add-apt-repository 'deb [arch=amd64] http://repo.sawtooth.me/ubuntu/bumper/stable xenial universe'
sudo apt-get update
sudo apt-get install -y sawtooth python3-sawtooth-poet-engine python3-sawtooth-identity sawtooth-devmode-engine-rust

# sudo -u sawtooth sawtooth-validator -vvv \
# --bind component:tcp://127.0.0.1:4004 \
# --bind network:tcp://172.31.82.76:8800 \
# --endpoint tcp://172.31.82.76:8800 \
# --peers tcp://172.31.81.27:8800,tcp://172.31.85.216

# sudo -u sawtooth sawtooth-validator -vvv \
# --bind component:tcp://127.0.0.1:4004 \
# --bind network:tcp://172.31.81.27:8800 \
# --endpoint tcp://172.31.81.27:8800 \
# --peers tcp://172.31.82.76:8800,tcp://172.31.85.216

# sudo -u sawtooth sawtooth-validator -vvv \
# --bind component:tcp://127.0.0.1:4004 \
# --bind network:tcp://172.31.85.216:8800 \
# --endpoint tcp://172.31.85.216:8800 \
# --peers tcp://172.31.82.76:8800,tcp://172.31.81.27:8800

# sawtooth keygen
# sudo sawadm keygen

# cd /tmp
# sudo -u sawtooth sawset genesis -k /etc/sawtooth/keys/validator.priv -o config-genesis.batch
# sudo -u sawtooth sawset proposal create -k /etc/sawtooth/keys/validator.priv \
#     -o config.batch \
#     sawtooth.consensus.algorithm=poet \
#     sawtooth.poet.report_public_key_pem="$(cat /etc/sawtooth/simulator_rk_pub.pem)" \
#     sawtooth.poet.valid_enclave_measurements=$(poet enclave measurement) \
#     sawtooth.poet.valid_enclave_basenames=$(poet enclave basename)
# sudo -u sawtooth poet registration create -k /etc/sawtooth/keys/validator.priv -o poet.batch

# # sudo -u sawtooth sawset proposal create -k /etc/sawtooth/keys/validator.priv \
# #     -o poet-settings.batch \
# #     sawtooth.poet.target_wait_time=5 \
# #     sawtooth.poet.initial_wait_time=25 \
# #     sawtooth.publisher.max_batches_per_block=100

# # sudo -u sawtooth sawadm genesis config-genesis.batch config.batch poet.batch poet-settings.batch

# sudo -u sawtooth sawtooth-validator \
#     --endpoint tcp://localhost:8800 \
#     --bind component:tcp://eth0:4004 \
#     --bind network:tcp://eth0:8800 \
#     --bind consensus:tcp://eth0:5050 \
#     -vvv


# sudo -u sawtooth sawtooth-validator \
#     --bind component:tcp://127.0.0.1:4004 \
#     --bind network:tcp://172.31.95.159:8800 \
#     --endpoint tcp://172.31.95.159:8800 \
#     --peers tcp://172.31.95.159:8800,tcp://172.31.84.149:8800


# sudo systemctl start sawtooth-rest-api.service
# sudo systemctl start sawtooth-poet-validator-registry-tp.service
# sudo systemctl start sawtooth-validator.service
# sudo systemctl start sawtooth-settings-tp.service
# sudo systemctl start sawtooth-intkey-tp-python.service
# sudo systemctl start sawtooth-identity-tp.service
# sudo systemctl start sawtooth-poet-engine.service

# sudo systemctl status sawtooth-rest-api.service
# sudo systemctl status sawtooth-poet-validator-registry-tp.service
# sudo systemctl status sawtooth-validator.service
# sudo systemctl status sawtooth-settings-tp.service
# sudo systemctl status sawtooth-intkey-tp-python.service
# sudo systemctl status sawtooth-identity-tp.service
# sudo systemctl status sawtooth-poet-engine.service

# sudo systemctl stop sawtooth-rest-api.service
# sudo systemctl stop sawtooth-poet-validator-registry-tp.service
# sudo systemctl stop sawtooth-validator.service
# sudo systemctl stop sawtooth-settings-tp.service
# sudo systemctl stop sawtooth-intkey-tp-python.service
# sudo systemctl stop sawtooth-identity-tp.service
# sudo systemctl stop sawtooth-poet-engine.service

# sawset genesis
# sudo -u sawtooth sawadm genesis config-genesis.batch
# sudo -u sawtooth sawtooth-validator -vv

# sudo -u sawtooth sawtooth-rest-api -v

# sudo rm -r /var/lib/sawtooth/
# sudo mkdir /var/lib/sawtooth
# sudo chmod -R 777 /var/lib/sawtooth
# sudo rm -r /var/log/sawtooth/
# sudo mkdir /var/log/sawtooth
# sudo chmod -R 777 /var/log/sawtooth

# sawset genesis
# sudo -u sawtooth sawadm genesis config-genesis.batch

# sudo rm /etc/sawtooth/keys/validator.*
# sudo rm /home/ubuntu/.sawtooth/keys/ubuntu.*

# docker
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update
sudo apt-get install -y docker-ce
sudo usermod -a -G docker ${USER}

# docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# go
curl -O https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.11.4.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
echo "export GOPATH=$HOME/go" >> ~/.profile
source ~/.profile
echo "export PATH=$PATH:$GOPATH/bin" >> ~/.profile
source ~/.profile
rm go1.11.4.linux-amd64.tar.gz
mkdir go

# nodejs
# curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -
# sudo apt-get install -y nodejs
# sudo apt-get install -y build-essential

# python 2.7
# sudo apt-get install -y python

# versions
# docker --version
# docker-compose --version
# go version
# echo "nodejs $(nodejs -v)"
# echo "npm $(npm -v)"
# python --version

sudo apt-get install -y libssl-dev
sudo apt-get install -y libzmq3-dev
sudo apt-get install -y build-essential
sudo apt-get install -y pkg-config
sudo apt-get install -y python3-pip
python3 -m pip install --upgrade pip
sudo python3 -m pip install grpcio
sudo python3 -m pip install grpcio-tools

go get -u \
    github.com/golang/protobuf/proto \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/pebbe/zmq4 \
    github.com/brianolson/cbor_go \
    github.com/satori/go.uuid \
    github.com/btcsuite/btcd/btcec \
    github.com/jessevdk/go-flags \
    github.com/pelletier/go-toml \
    github.com/golang/mock/gomock \
    github.com/golang/mock/mockgen \
    golang.org/x/crypto/ripemd160 \
    golang.org/x/crypto/ssh \
    github.com/hyperledger/sawtooth-sdk-go \
    github.com/decred/dcrd/dcrec/secp256k1 \
    github.com/gin-gonic/gin

cd $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go
go generate
