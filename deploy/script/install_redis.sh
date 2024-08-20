#!/bin/bash
echo "=============== start build redis ========================="

# 安装 Redis
apt-get update
apt-get install -y redis-server

# 修改 Redis 配置文件，确保监听所有接口
sed -i 's/^bind 127.0.0.1 ::1/bind 0.0.0.0/' /etc/redis/redis.conf
sed -i 's/^protected-mode yes/protected-mode no/' /etc/redis/redis.conf

# 启动 Redis 服务
redis-server /etc/redis/redis.conf
