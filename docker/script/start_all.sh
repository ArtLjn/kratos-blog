#!/bin/bash
time_now=$(date "+%Y-%m-%d %H:%M:%S")
echo -e "\033[32m
    ==========================================
    当前时间: $time_now
    功能：服务一键启动
    作者: @Ljn
    官网: https://lllcnm.cn
    ==========================================
    \033[0m"

# 设置公共路径
os_path="/usr/local/hongDou/"

function Run() {
    # 使用公共路径来启动各服务并重定向输出到对应日志文件
    nohup ${os_path}blog/blog >> ${os_path}log/blog.log 2>&1 &
    sleep 1
    nohup ${os_path}comment/comment >> ${os_path}log/comment.log 2>&1 &
    sleep 1
    nohup ${os_path}user/user >> ${os_path}log/user.log 2>&1 &
    sleep 1
    nohup ${os_path}gateway/gateway >> ${os_path}log/gateway.log 2>&1 &
}

function CheckProcess() {
    sleep 3
    # 检查服务进程并输出到process.log文件
    ps -ef | grep blog >> ${os_path}log/process.log
    ps -ef | grep comment >> ${os_path}log/process.log
    ps -ef | grep user >> ${os_path}log/process.log
    ps -ef | grep gateway >> ${os_path}log/process.log
}

function CheckPort() {
    sleep 3
    # 检查端口
    netstat -tuln | grep 8000 > process.log
    netstat -tuln | grep 8001 > process.log
    netstat -tuln | grep 8002 > process.log
    netstat -tuln | grep 8080 > process.log
}

Run
CheckProcess
CheckPort