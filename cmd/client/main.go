package main

import (
	"context"
	"fmt"
	"github.com/0xweb-3/temporal_example/internal/workflows"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

const TemporalUrl = "192.168.21.4:7233"

func main() {
	// 连接 Temporal Server
	c, err := client.Dial(client.Options{
		HostPort: TemporalUrl,
	})
	if err != nil {
		panic("无法连接 Temporal")
	}
	defer c.Close()

	// 启动 Workflow
	workflowOptions := client.StartWorkflowOptions{
		//ID: "order_workflow_123",
		ID:        "order_" + uuid.NewString(),
		TaskQueue: "ORDER_TASK_QUEUE",
	}

	// order123 假装被当作订单id传递，也就是说这里是一些业务id
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.OrderWorkflow, "order123")
	if err != nil {
		panic("无法启动 Workflow")
	}

	fmt.Println("订单处理 Workflow 启动成功，WorkflowID:", we.GetID())

	// 等待 Workflow 结束
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		panic("Workflow 执行失败")
	}

	fmt.Println("订单处理成功，结果:", result)
}
