#!/bin/bash
sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
apt update
apt-get install -y curl unzip wget
source ./install_consul.sh
source ./install_mysql.sh
source ./install_rabbitmq.sh
source ./install_redis.sh


