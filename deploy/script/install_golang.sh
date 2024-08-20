!/bin/bash
# 下载 Go 语言的最新版本
GO_VERSION="1.20.5"
wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz

# 解压并安装
tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz

# 设置 Go 环境变量和代理
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
echo "export GOPROXY=https://goproxy.cn,direct" >> ~/.profile
echo "export GO111MODULE=on" >> ~/.profile
source ~/.profile