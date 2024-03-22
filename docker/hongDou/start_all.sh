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
      echo "---------------sleep 30--------------------------"
      sleep 30
      echo "-----------------sleep end -----------------------"
      nohup ${os_path}tool/tool >> ${os_path}log/tool.log 2>&1 &
      nohup ${os_path}blog/blog >> ${os_path}log/blog.log 2>&1 &
      sleep 5
      nohup ${os_path}comment/comment >> ${os_path}log/comment.log 2>&1 &
      sleep 5
      nohup ${os_path}user/user >> ${os_path}log/user.log 2>&1 &
      sleep 5
      ./gateway/gateway
}

Run