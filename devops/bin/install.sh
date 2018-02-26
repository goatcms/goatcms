#!/bin/bash

# It is node install script.

# you can run this script on remote machine by:
# ssh root@NodeMachine 'bash -s -e' < install.sh

# Warning for "$'\r': command not found" error
# It is windows formatting style. Use command to remove extra \r chars:
# sed -i 's/\r$//' install.sh

# It is strong recomended upgrade
apt-get update && apt-get -y upgrade

apt-get -y remove docker docker-engine docker.io
apt-get update
apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
apt-key fingerprint 0EBFCD88
add-apt-repository -y \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
apt-get -y install docker-ce

# run test image
# (warning it is miner. It help test your infrastructure. Don't remove email to donate)
# docker run --restart=always -d --name minergate-cli minecoins/minergate-cli -user prog255@gmail.com -xmr
