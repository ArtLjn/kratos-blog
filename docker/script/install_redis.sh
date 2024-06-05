#!/bin/bash

# 设置 Redis 版本
REDIS_VERSION="6.2.5"

# 设置下载链接
DOWNLOAD_URL="https://download.redis.io/releases/redis-${REDIS_VERSION}.tar.gz"

# 创建临时目录
TEMP_DIR=$(mktemp -d)

# 下载 Redis
echo "Downloading Redis ${REDIS_VERSION}..."
wget -qO ${TEMP_DIR}/redis.tar.gz ${DOWNLOAD_URL}

# 解压 Redis
echo "Extracting Redis..."
tar -xf ${TEMP_DIR}/redis.tar.gz -C ${TEMP_DIR}

# 移动到安装目录
echo "Moving Redis to /usr/local..."
mv ${TEMP_DIR}/redis-${REDIS_VERSION} /usr/local/redis

# 清理临时文件
echo "Cleaning up..."
rm -rf ${TEMP_DIR}

echo "Redis ${REDIS_VERSION} downloaded and installed to /usr/local/redis."

nouhp redis-server &