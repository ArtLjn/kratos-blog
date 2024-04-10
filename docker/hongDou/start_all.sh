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

function Run() {
      echo "---------------sleep 30--------------------------"
      sleep 30
      echo "-----------------sleep end -----------------------"
      nohup tool/tool >> log/tool.log 2>&1 &
      nohup blog/blog >> log/blog.log 2>&1 &
      sleep 5
      nohup comment/comment >> log/comment.log 2>&1 &
      sleep 5
      nohup user/user >> log/user.log 2>&1 &
      sleep 5
      ./gateway/gateway
}

Run