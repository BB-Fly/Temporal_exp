#!/bin/bash

# 启动脚本，用于启动temporal server、worker和工作流

# 检查temporal是否安装
if ! command -v temporal &> /dev/null; then
    echo "错误: temporal 命令未找到，请先安装 temporal CLI"
    exit 1
fi

# 检查go是否安装
if ! command -v go &> /dev/null; then
    echo "错误: go 命令未找到，请先安装 Go"
    exit 1
fi

# 定义颜色
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
RED="\033[0;31m"
NC="\033[0m" # No Color

echo -e "${GREEN}=== 开始启动 Temporal 工作流系统 ===${NC}"

# 启动 temporal server
echo -e "${YELLOW}1. 启动 Temporal Server...${NC}"
temporal server start-dev > temporal-server.log 2>&1 &
SERVER_PID=$!
echo -e "${GREEN}Temporal Server 已启动，进程ID: ${SERVER_PID}${NC}"

# 等待服务器启动
 sleep 5

echo -e "${YELLOW}2. 启动 Worker...${NC}"
# 启动 worker
go run src/worker/main.go > worker.log 2>&1 &
WORKER_PID=$!
echo -e "${GREEN}Worker 已启动，进程ID: ${WORKER_PID}${NC}"

# 等待worker启动
 sleep 2

echo -e "${YELLOW}3. 启动工作流...${NC}"
# 启动批量工作流
go run src/start/main.go batch > workflow.log 2>&1 &
WORKFLOW_PID=$!
echo -e "${GREEN}工作流已启动，进程ID: ${WORKFLOW_PID}${NC}"

# 等待工作流启动
 sleep 2

echo -e "${YELLOW}4. 打开 Temporal UI 页面...${NC}"
# 打开 Temporal UI 页面
if command -v open &> /dev/null; then
    open http://localhost:8233
elif command -v xdg-open &> /dev/null; then
    xdg-open http://localhost:8233
else
    echo -e "${YELLOW}请手动打开浏览器访问: http://localhost:8233${NC}"
fi

echo -e "${GREEN}=== 启动完成 ===${NC}"
echo -e "${YELLOW}查看日志:${NC}"
echo -e "  - Temporal Server: tail -f temporal-server.log"
echo -e "  - Worker: tail -f worker.log"
echo -e "  - 工作流: tail -f workflow.log"
echo -e "${YELLOW}停止所有进程:${NC}"
echo -e "  kill ${SERVER_PID} ${WORKER_PID} ${WORKFLOW_PID}"
