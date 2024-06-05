#!/bin/bash

time_now=$(date "+%Y-%m-%d %H:%M:%S")
blogServerPortList=(9000 9001 9002)
toolPort=8099
gatewayPort=8080

echo -e "\033[32m
    ==========================================
    当前时间: $time_now
    功能：服务一键启动
    作者: @Ljn
    官网: https://lllcnm.cn
    ==========================================
    \033[0m"

function Run() {
    echo "---------------sleep 30--------------------------"
    sleep 30
    echo "-----------------sleep end -----------------------"
    nohup tool/tool &
    nohup blog/blog &
    nohup comment/comment &
    nohup user/user &
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
    nohup gateway/gateway &
}

Run
CheckBlogServerRunningStatus
CheckToolRunningStatus
StartGateway

echo "所有服务已成功启动。"
