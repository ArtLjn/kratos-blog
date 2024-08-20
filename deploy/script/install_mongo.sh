#!/bin/bash
echo "=============== start build mongo ========================="

# 安装必要的依赖
apt-get install -y gnupg wget

# 导入MongoDB GPG密钥
wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | apt-key add -

# 创建MongoDB源列表
echo "deb [ arch=amd64,arm64 ] http://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-6.0.list

# 更新包索引并安装MongoDB
apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -y mongodb-org --fix-missing

mkdir -p /data/db

nohup mongod --dbpath /data/db > /dev/null 2>&1 &
