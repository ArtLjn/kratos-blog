#!/bin/bash
echo "=============== start build rabbitmq ========================="

set -e

# 更新包列表并安装 Erlang
echo "更新包列表并安装 Erlang..."
DEBIAN_FRONTEND=noninteractive apt-get -y -q install erlang

# 安装 RabbitMQ
echo "安装 RabbitMQ..."
apt-get install -y rabbitmq-server

# 设置环境变量
export PATH=$PATH:/usr/sbin:/usr/bin:/sbin:/bin

# 启动 RabbitMQ
echo "启动 RabbitMQ..."
rabbitmq-server -detached

# 启用管理插件
echo "启用 RabbitMQ 管理插件..."
rabbitmq-plugins enable rabbitmq_management

# 重启 RabbitMQ 以确保插件生效
echo "重启 RabbitMQ..."
rabbitmqctl stop
sleep 5
rabbitmq-server -detached

