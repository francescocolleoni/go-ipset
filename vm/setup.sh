#!/bin/sh

apt-get update
apt-get install -y ipset

echo "downloading go package (1.18)..."
wget https://go.dev/dl/go1.18.linux-amd64.tar.gz -O gopkg.tar.gz

echo "download complete, unpacking..."
rm -rf /usr/local/go && tar -C /usr/local -xzf gopkg.tar.gz
rm gopkg.tar.gz

echo "adding entry to /etc/profile..."
printf "\nexport PATH=$PATH:/usr/local/go/bin" >> /etc/profile
printf "\nexport PATH=$PATH:/usr/local/go/bin" >> ~/.profile

. /etc/profile
. ~/.profile

go version
