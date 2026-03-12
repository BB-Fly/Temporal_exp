#!/bin/bash

# 控制脚本，用于启动、停止服务，以及启动不同的工作流

# 检查temporal是否安装
if ! command -v temporal &> /dev/null; then
    echo "错误: temporal 命令未找到，请先安装 temporal CLI"
    exit 1
fi

# 检查编译后的文件是否存在
if [ ! -f "bin/worker" ] || [ ! -f "bin/start" ]; then
    echo "错误: 可执行文件不存在，请先运行 build.sh 编译项目"
    exit 1
fi

# 定义颜色
GREEN="\033[0;32m"
YELLOW="\033[1;33m"
RED="\033[0;31m"
NC="\033[0m" # No Color

# 定义日志文件
SERVER_LOG="temporal-server.log"
WORKER_LOG="worker.log"

# 启动服务
start() {
    echo -e "${GREEN}=== 开始启动 Temporal 工作流系统 ===${NC}"
    
    # 启动 temporal server
    echo -e "${YELLOW}1. 启动 Temporal Server...${NC}"
    temporal server start-dev > $SERVER_LOG 2>&1 &
    SERVER_PID=$!
    echo -e "${GREEN}Temporal Server 已启动，进程ID: ${SERVER_PID}${NC}"
    
    # 等待服务器启动
    sleep 5
    
    echo -e "${YELLOW}2. 启动 Worker...${NC}"
    # 启动 worker
    bin/worker > $WORKER_LOG 2>&1 &
    WORKER_PID=$!
    echo -e "${GREEN}Worker 已启动，进程ID: ${WORKER_PID}${NC}"
    
    # 等待worker启动
    sleep 2
    
    # 保存进程ID到文件
    echo $SERVER_PID > server.pid
    echo $WORKER_PID > worker.pid
    
    echo -e "${GREEN}=== 服务启动完成 ===${NC}"
    echo -e "${YELLOW}查看日志:${NC}"
    echo -e "  - Temporal Server: tail -f $SERVER_LOG"
    echo -e "  - Worker: tail -f $WORKER_LOG"
}

# 停止服务
stop() {
    echo -e "${GREEN}=== 开始停止 Temporal 工作流系统 ===${NC}"
    
    # 停止工作流进程
    if [ -f "workflow.pid" ]; then
        WORKFLOW_PIDs=$(cat workflow.pid)
        for PID in $WORKFLOW_PIDs; do
            if ps -p $PID > /dev/null; then
                echo -e "${YELLOW}停止工作流进程: $PID${NC}"
                kill $PID
            fi
        done
        rm -f workflow.pid
    fi
    
    # 停止worker
    if [ -f "worker.pid" ]; then
        WORKER_PID=$(cat worker.pid)
        if ps -p $WORKER_PID > /dev/null; then
            echo -e "${YELLOW}停止 Worker 进程: $WORKER_PID${NC}"
            kill $WORKER_PID
        fi
        rm -f worker.pid
    fi
    
    # 停止temporal server
    if [ -f "server.pid" ]; then
        SERVER_PID=$(cat server.pid)
        if ps -p $SERVER_PID > /dev/null; then
            echo -e "${YELLOW}停止 Temporal Server 进程: $SERVER_PID${NC}"
            kill $SERVER_PID
        fi
        rm -f server.pid
    fi
    
    echo -e "${GREEN}=== 服务停止完成 ===${NC}"
}

# 启动工作流
start_workflow() {
    WORKFLOW_TYPE=$1
    shift
    
    echo -e "${GREEN}=== 启动工作流: $WORKFLOW_TYPE ===${NC}"
    
    # 检查服务是否已启动
    if [ ! -f "server.pid" ] || [ ! -f "worker.pid" ]; then
        echo -e "${YELLOW}服务未启动，正在启动服务...${NC}"
        start
    fi
    
    # 启动工作流
    bin/start $WORKFLOW_TYPE "$@" > workflow-${WORKFLOW_TYPE}.log 2>&1 &
    WORKFLOW_PID=$!
    echo -e "${GREEN}工作流 $WORKFLOW_TYPE 已启动，进程ID: ${WORKFLOW_PID}${NC}"
    
    # 保存工作流进程ID
    if [ -f "workflow.pid" ]; then
        echo $WORKFLOW_PID >> workflow.pid
    else
        echo $WORKFLOW_PID > workflow.pid
    fi
    
    echo -e "${GREEN}=== 工作流启动完成 ===${NC}"
    echo -e "${YELLOW}查看日志: tail -f workflow-${WORKFLOW_TYPE}.log${NC}"
}

# 启动多个工作流
start_multiple_workflows() {
    echo -e "${GREEN}=== 启动多个工作流 ===${NC}"
    
    # 检查服务是否已启动
    if [ ! -f "server.pid" ] || [ ! -f "worker.pid" ]; then
        echo -e "${YELLOW}服务未启动，正在启动服务...${NC}"
        start
    fi
    
    # 启动所有指定的工作流
    for WORKFLOW_TYPE in "$@"; do
        echo -e "${YELLOW}启动工作流: $WORKFLOW_TYPE${NC}"
        bin/start $WORKFLOW_TYPE > workflow-${WORKFLOW_TYPE}.log 2>&1 &
        WORKFLOW_PID=$!
        echo -e "${GREEN}工作流 $WORKFLOW_TYPE 已启动，进程ID: ${WORKFLOW_PID}${NC}"
        
        # 保存工作流进程ID
        if [ -f "workflow.pid" ]; then
            echo $WORKFLOW_PID >> workflow.pid
        else
            echo $WORKFLOW_PID > workflow.pid
        fi
        
        # 等待短暂时间，避免同时启动过多工作流
        sleep 1
    done
    
    echo -e "${GREEN}=== 多个工作流启动完成 ===${NC}"
    echo -e "${YELLOW}查看日志: tail -f workflow-*.log${NC}"
}

# 打开 Temporal UI
open_ui() {
    echo -e "${GREEN}=== 打开 Temporal UI ===${NC}"
    if command -v open &> /dev/null; then
        open http://localhost:8233
    elif command -v xdg-open &> /dev/null; then
        xdg-open http://localhost:8233
    else
        echo -e "${YELLOW}请手动打开浏览器访问: http://localhost:8233${NC}"
    fi
}

# 显示帮助信息
show_help() {
    echo "使用方法: $0 [命令] [参数]"
    echo ""
    echo "命令:"
    echo "  start                启动 Temporal 服务（包括 server 和 worker）"
    echo "  stop                 停止 Temporal 服务和所有工作流"
    echo "  workflow [类型]      启动指定类型的工作流"
    echo "  multiple [类型1] [类型2] ... 启动多个工作流"
    echo "  ui                   打开 Temporal UI 页面"
    echo "  help                 显示此帮助信息"
    echo ""
    echo "工作流类型:"
    echo "  greeting             启动问候工作流"
    echo "  schedule             启动班次推荐工作流"
    echo "  batch                批量启动班次推荐工作流"
    echo "  prelock              启动预锁库存工作流"
    echo ""
    echo "示例:"
    echo "  $0 start                     # 启动服务"
    echo "  $0 workflow batch             # 启动批量工作流"
    echo "  $0 multiple schedule prelock  # 同时启动班次推荐和预锁库存工作流"
    echo "  $0 stop                      # 停止所有服务和工作流"
}

# 主函数
main() {
    if [ $# -eq 0 ]; then
        show_help
        exit 1
    fi
    
    case $1 in
        start)
            start
            ;;
        stop)
            stop
            ;;
        workflow)
            if [ $# -lt 2 ]; then
                echo -e "${RED}错误: 请指定工作流类型${NC}"
                show_help
                exit 1
            fi
            shift
            start_workflow "$@"
            ;;
        multiple)
            if [ $# -lt 2 ]; then
                echo -e "${RED}错误: 请至少指定一个工作流类型${NC}"
                show_help
                exit 1
            fi
            shift
            start_multiple_workflows "$@"
            ;;
        ui)
            open_ui
            ;;
        help)
            show_help
            ;;
        *)
            echo -e "${RED}错误: 未知命令 '$1'${NC}"
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"
