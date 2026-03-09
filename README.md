# Temporal Go 示例项目

这是一个使用 Temporal Go SDK 构建的简单示例项目，演示了如何创建和运行一个基本的工作流。

## 环境依赖
首先，按照官方文档安装环境依赖
- [Temporal Go SDK 文档](https://docs.temporal.io/docs/go/)

具体包括：
- 安装并检查go环境：go version
- 安装并检查brew环境：brew --version
- 安装temporal：brew install temporal
- 启动temporal server：temporal server start-dev

## 项目结构

```
Temporal_exp/
├── src/
│   ├── greating/                 # HelloWorld工作流模块
│   │   ├── activity              # 活动定义
│   │   └── workflow              # 工作流定义
│   ├── start                     # 工作流启动器
│   └── worker                    # Worker 启动器
├── vendor/                       # 依赖包
├── go.mod                        # Go 模块文件
├── go.sum                        # 依赖校验和
└── README.md                     # 项目说明
```

## 项目功能

该项目实现了若干Temporal工作流示例，主要包含以下组件：

1. **活动 (Activity)**：Activityle 定义的活动。
2. **工作流 (Workflow)**：Workflow 类定义工作流。
3. **Worker**：注册工作流和活动，处理任务队列中的任务
4. **启动器**：启动工作流实例，传递参数并获取结果

## 运行程序

### 前提条件

- 已安装 Go 环境
- Temporal 服务器正在运行（默认监听 localhost:7233）

### 1. 启动 Worker（终端 1）

```bash
go run src/worker/main.go
```

### 2. 启动工作流（终端 2）

```bash
go run src/start/main.go <name>
```

例如：

```bash
go run src/start/main.go World
```

### 查看结果

成功运行后，终端 2 会输出类似以下内容：

```
2026/03/09 10:00:00 Started workflow WorkflowID greeting-workflow RunID <随机ID>
2026/03/09 10:00:00 Workflow result: Hello World
```

## 在 Web UI 中查看

打开浏览器访问 Temporal Web UI，您可以看到工作流的执行历史和事件：

http://localhost:8233 → 查看工作流执行历史

## 代码说明-以HelloWorld工作流为例

### 活动定义 (`src/greating/activity/a_greeting.go`)

```go
func Greet(ctx context.Context, name string) (string, error) {
    return fmt.Sprintf("Hello %s", name), nil
}
```

### 工作流定义 (`src/greating/workflow/wf_greating.go`)

```go
func SayHelloWorkflow(ctx workflow.Context, name string) (string, error) {
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 10,
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    var result string
    err := workflow.ExecuteActivity(ctx, greeting.Greet, name).Get(ctx, &result)
    if err != nil {
        return "", err
    }

    return result, nil
}
```

### Worker 启动 (`src/worker/main.go`)

注册工作流和活动，并启动 Worker 处理任务队列。

### 工作流启动 (`src/start/main.go`)

创建 Temporal 客户端，启动工作流实例，并获取执行结果。

## 参考文档

- [Temporal Go SDK 文档](https://docs.temporal.io/docs/go/)
- [Temporal 官方文档](https://docs.temporal.io/)