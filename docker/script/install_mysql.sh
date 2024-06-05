#!/bin/bash

# 设置 MySQL 版本
MYSQL_VERSION="8.0.27-1ubuntu20.04"

# 设置 MYSQL 密码
MYSQL_ROOT_PASSWORD="123456"

# 设置下载链接（使用阿里云的镜像源）
DOWNLOAD_URL="https://mirrors.aliyun.com/mysql/MySQL-8.0/mysql-server_${MYSQL_VERSION}_amd64.deb-bundle.tar"

# 创建临时目录
TEMP_DIR=$(mktemp -d)

# 下载 MySQL
echo "Downloading MySQL ${MYSQL_VERSION} from Aliyun Mirror..."
#wget -O ${TEMP_DIR}/mysql-server.deb-bundle.tar ${DOWNLOAD_URL}

# 检查下载是否成功
#if [ $? -ne 0 ]; then
#    echo "Failed to download MySQL from ${DOWNLOAD_URL}"
#    exit 1
#fi
cp -r  mysql-server_8.0.27-1ubuntu20.04_amd64.deb-bundle.tar ${TEMP_DIR}/mysql-server.deb-bundle.tar
# 解压 MySQL
echo "Extracting MySQL..."
tar -xf ${TEMP_DIR}/mysql-server.deb-bundle.tar -C ${TEMP_DIR}

# 检查解压是否成功
if [ $? -ne 0 ]; then
    echo "Failed to extract MySQL tar file"
    exit 1
fi

# 配置预先设置文件，以避免在安装过程中出现交互提示
echo "Configuring preseed file for MySQL installation..."
echo "mysql-community-server mysql-community-server/root-pass password ${MYSQL_ROOT_PASSWORD}" |  debconf-set-selections
echo "mysql-community-server mysql-community-server/re-root-pass password ${MYSQL_ROOT_PASSWORD}" |  debconf-set-selections

# 安装 MySQL
echo "Installing MySQL..."
dpkg -i ${TEMP_DIR}/mysql-common_${MYSQL_VERSION}_amd64.deb
dpkg -i ${TEMP_DIR}/mysql-community-client-plugins_${MYSQL_VERSION}_amd64.deb
dpkg -i ${TEMP_DIR}/mysql-community-client-core_${MYSQL_VERSION}_amd64.deb
dpkg -i ${TEMP_DIR}/mysql-community-client_${MYSQL_VERSION}_amd64.deb
dpkg -i ${TEMP_DIR}/mysql-community-server-core_${MYSQL_VERSION}_amd64.deb
dpkg -i ${TEMP_DIR}/mysql-community-server_${MYSQL_VERSION}_amd64.deb

# 检查安装是否成功
if [ $? -ne 0 ]; then
    echo "Failed to install MySQL"
    exit 1
fi

# 清理临时文件
echo "Cleaning up..."
#rm -rf ${TEMP_DIR}

echo "MySQL ${MYSQL_VERSION} downloaded and installed with default password '${MYSQL_ROOT_PASSWORD}'."
