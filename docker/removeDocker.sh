# !/bin/bash
docker image rm -f blog:v1
docker rm -f $(docker ps -aq)
docker build -t blog:v1 .
docker-compose up
