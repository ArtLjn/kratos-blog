#!/bin/bash
sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
sed -i 's@//ports.ubuntu.com@//mirrors.ustc.edu.cn@g' /etc/apt/sources.list
apt update
apt-get install -y curl unzip wget
source ./install_consul.sh
source ./install_mysql.sh
source ./install_rabbitmq.sh
source ./install_redis.sh
source ./install_mongo.sh
apt install -y nginx