#!/bin/bash

time_now=$(date "+%Y-%m-%d %H:%M:%S")
blogServerPortList=(9000 9001 9002)
toolPort=8099
gatewayPort=8080
FLUSH PRIVILEGES;

echo -e "\033[32m
    ==========================================
    当前时间: $time_now
    功能：服务一键启动
    作者: @Ljn
    官网: https://lllcnm.cn
    ==========================================
    \033[0m"

function Run() {
    rabbitmq-server -detached
    # 等待 RabbitMQ 启动
    sleep 5
    # 启动 MySQL 服务
    service mysql start
    # 启动 Consul
    nohup consul agent -dev -client=0.0.0.0 -config-dir=/etc/consul.d >/dev/null 2>&1 &
    # 启动 Redis
    nohup redis-server >/dev/null 2>&1 &
    echo "---------------start build--------------------------"
    # 等待一段时间
    sleep 10
    # 确保 RabbitMQ 应用程序正在运行
    echo "确保 RabbitMQ 应用程序正在运行..."
    rabbitmqctl start_app

    echo "Erlang 和 RabbitMQ 安装完成，RabbitMQ 已启动并启用管理插件。"

    # 添加管理员用户
    echo "添加管理员用户 root..."
    rabbitmqctl add_user root 123456
    rabbitmqctl set_user_tags root administrator

    # 创建虚拟主机
    echo "创建虚拟主机 hongDou..."
    rabbitmqctl add_vhost hongDou

    # 为 root 用户设置虚拟主机权限
    echo "为 root 用户设置虚拟主机 hongDou 的权限..."
    rabbitmqctl set_permissions -p hongDou root ".*" ".*" ".*"

    echo "RabbitMQ 安装和配置完成。"
    # 启动 MongoDB
    nohup mongod --dbpath /data/db > /dev/null 2>&1 &
    # 安装 lsof
    apt install -y lsof
    # 启动其他服务
    nohup tool/tool >/dev/null 2>&1 &
    nohup blog/blog >/dev/null 2>&1 &
    nohup comment/comment >/dev/null 2>&1 &
    nohup user/user >/dev/null 2>&1 &
    nginx
}

function CheckPortRunningStatus() {
    local port=$1
    if lsof -i:$port >/dev/null 2>&1; then
        return 0
    else
        return 1
    fi
}

function CheckBlogServerRunningStatus() {
    for port in "${blogServerPortList[@]}"; do
        while ! CheckPortRunningStatus $port; do
            echo "端口 $port 上的博客服务器未运行。5秒后再检查..."
            sleep 5
        done
        echo "端口 $port 上的博客服务器正在运行。"
    done
}

function CheckToolRunningStatus() {
    while ! CheckPortRunningStatus $toolPort; do
        echo "端口 $toolPort 上的工具服务未运行。5秒后再检查..."
        sleep 5
    done
    echo "端口 $toolPort 上的工具服务正在运行。"
}

function StartGateway() {
    echo "正在启动端口 $gatewayPort 上的网关服务..."
    ./gateway/gateway
}

Run
CheckBlogServerRunningStatus
CheckToolRunningStatus
StartGateway

echo "所有服务已成功启动。"
