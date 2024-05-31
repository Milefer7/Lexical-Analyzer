#!/bin/bash

# 进入后端目录并在后台启动后端服务
cd be && go run main.go &

# 输出提示信息
echo "后端服务正在启动..."

# 等待一段时间以确保后端服务已经启动
sleep 5

# 进入前端目录并启动前端服务
echo "前端服务正在启动..."
cd fe && npm run dev