#!/bin/bash

# 更新包列表
apt-get update

# 安装 MySQL Server
DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-server
DEBIAN_FRONTEND=noninteractive apt-get install -y mysql-client


# 启动 MySQL 服务
service mysql start

# 确保 MySQL 服务已启动
sleep 5

# 进行 MySQL 安全设置和密码更新
mysql -u root <<EOF
INSTALL PLUGIN validate_password SONAME 'validate_password.so';
SET GLOBAL validate_password_policy=0;
SET GLOBAL validate_password_mixed_case_count=0;
SET GLOBAL validate_password_number_count=0;
SET GLOBAL validate_password_special_char_count=0;
SET GLOBAL validate_password_length=6;
SHOW VARIABLES LIKE 'validate_password%';
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
UPDATE mysql.user SET Host='%' WHERE User='root' AND Host='localhost';FLUSH PRIVILEGES;
EOF

# 创建 hongDou 数据库
mysql -u root -p123456 -e "CREATE DATABASE IF NOT EXISTS hongDou CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"

# 导入 SQL 文件到 hongDou 数据库
mysql -u root -p123456 hongDou < ../hongDou/sql/blog.sql
sed -i 's/^bind-address\s*=\s*127.0.0.1/bind-address = 0.0.0.0/' /etc/mysql/mysql.conf.d/mysqld.cnf
service mysql restart
# 输出 MySQL 状态
service mysql status
