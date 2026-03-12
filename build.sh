#!/bin/bash

# 编译脚本，用于编译项目的worker和start组件

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

echo -e "${GREEN}=== 开始编译 Temporal 项目 ===${NC}"

# 创建输出目录
mkdir -p bin

# 编译 worker
echo -e "${YELLOW}1. 编译 Worker...${NC}"
go build -o bin/worker src/worker/main.go
if [ $? -ne 0 ]; then
    echo -e "${RED}Worker 编译失败${NC}"
    exit 1
fi
echo -e "${GREEN}Worker 编译成功${NC}"

# 编译 start
echo -e "${YELLOW}2. 编译 Start...${NC}"
go build -o bin/start src/start/main.go
if [ $? -ne 0 ]; then
    echo -e "${RED}Start 编译失败${NC}"
    exit 1
fi
echo -e "${GREEN}Start 编译成功${NC}"

echo -e "${GREEN}=== 编译完成 ===${NC}"
echo -e "编译结果位于: ${YELLOW}bin/ 目录${NC}"
echo -e "可执行文件:"
echo -e "  - bin/worker: 工作流 worker"
echo -e "  - bin/start: 工作流启动器"
