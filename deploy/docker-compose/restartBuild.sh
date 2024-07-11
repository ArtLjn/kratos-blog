# !/bin/bash
docker image rm -f blog:v1
docker rm -f hongDou_blog hongDou_mysql hongDou_redis hongDou_consul
docker build -t blog:v1 .
docker-compose up
