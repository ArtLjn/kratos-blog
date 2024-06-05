#!/bin/bash
# 安装 Erlang
echo "Installing Erlang..."
DEBIAN_FRONTEND=noninteractive apt-get -y -q install erlang

# 安装 RabbitMQ
echo "Installing RabbitMQ..."
apt-get install -y rabbitmq-server

# 设置环境变量
export PATH=$PATH:/usr/sbin:/usr/bin:/sbin:/bin

# 启动 RabbitMQ
echo "Starting RabbitMQ..."
rabbitmq-server -detached

echo "Erlang and RabbitMQ installed and RabbitMQ started."
