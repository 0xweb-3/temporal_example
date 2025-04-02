# temporal golang案例
## 设计的temporal使用背景
我们要构建一个 订单处理系统，当用户下单后，系统会执行以下步骤：
1. 验证库存 
2. 扣减库存 
3. 处理支付 
4. 通知用户订单成功

**整个过程中可能的问题**
* 失败（如库存不足、支付失败）
* 需要重试（如支付超时）
* 需要异步处理（如支付成功后通知用户）

**Temporal在其中的作用**
* 每个步骤都会执行，即使服务器崩溃也不会丢失任务。
* 失败的任务会自动重试，避免因瞬时故障导致任务失败。

## 安装依赖
```shell
go mod init temporal-example
go get go.temporal.io/sdk
```

## 目录结构说明
在 **Golang** 项目中，按照最佳实践，我们可以使用以下目录结构来组织 **Temporal 订单处理系统**：

```
temporal-order-system/
│── go.mod
│── go.sum
│── README.md
│
├── cmd/
│   ├── client/            # 触发 Workflow 的客户端
│   │   ├── main.go        # 运行 Client 启动 Workflow
│   │
│   ├── worker/            # Worker 进程，负责执行 Workflow 和 Activity
│   │   ├── main.go        # 运行 Worker 监听任务队列
│
├── internal/
│   ├── workflows/         # 业务流程（Workflow）
│   │   ├── order_workflow.go  # 订单处理 Workflow 逻辑
│   │
│   ├── activities/        # 具体业务逻辑（Activity）
│   │   ├── stock.go       # 库存相关 Activity
│   │   ├── payment.go     # 支付相关 Activity
│   │   ├── notification.go # 通知用户的 Activity
│
└── docker-compose.yml     # 启动 Temporal Server 的 Docker 配置（可选）
```
---

## **📌 运行步骤**
### **1️⃣ 启动 Temporal Server**
如果你使用 Docker：
```sh
docker-compose up -d
```
或者使用 `temporalite`：
```sh
temporalite start
```

### **2️⃣ 启动 Worker**
```sh
go run cmd/worker/main.go
```

### **3️⃣ 运行 Client 触发 Workflow**
```sh
go run cmd/client/main.go
```

---

## **📌 总结**
✅ **代码模块化**，易维护  
✅ **Activity 独立拆分**，业务清晰  
✅ **Temporal 提供任务可靠性**，支持失败重试

**这样，订单处理系统就可以高可靠、高扩展地运行了！** 🎯🚀


## ActivityOptions参数
* StartToCloseTimeout	Activity从开始到完成的最长时间
* ScheduleToCloseTimeout 从Activity被调度到完成的最长时间（包括排队等待的时间）
* ScheduleToStartTimeout 从Activity被调度到开始执行的最长时间（Worker太忙可能导致排队）
* HeartbeatTimeout 定期心跳超时时间，用于长时间运行的Activity防止中断
```go
// 重试策略
ao := workflow.ActivityOptions{
    StartToCloseTimeout: time.Second * 10,
    RetryPolicy: &temporal.RetryPolicy{
        InitialInterval:    time.Second * 1,  // 第一次重试间隔
        BackoffCoefficient: 2.0,              // 指数级退避策略（翻倍）
        MaximumInterval:    time.Second * 10, // 最大重试间隔
        MaximumAttempts:    5,                // 最多重试 5 次
    },
}
```