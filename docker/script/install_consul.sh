#!/bin/bash

# 设置 Consul 版本
CONSUL_VERSION="1.13.1"

# 下载 Consul
curl -O https://releases.hashicorp.com/consul/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_linux_amd64.zip

# 解压并安装 Consul
unzip consul_${CONSUL_VERSION}_linux_amd64.zip
mv consul /usr/local/bin/

# 验证安装
consul --version

# 创建 Consul 配置目录和数据目录
mkdir -p /etc/consul.d
mkdir -p /var/lib/consul

# 创建默认配置文件
cat <<EOF > /etc/consul.d/consul.hcl
data_dir = "/var/lib/consul"
EOF

# 后台启动 Consul
nohup consul agent -dev -client=0.0.0.0 -config-dir=/etc/consul.d &

echo "Consul 启动完成。"
