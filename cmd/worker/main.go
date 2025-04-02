package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"

	"github.com/0xweb-3/temporal_example/internal/activities"
	"github.com/0xweb-3/temporal_example/internal/workflows"
)

const TemporalUrl = "192.168.21.4:7233"

func main() {
	// 连接 Temporal Server
	c, err := client.Dial(client.Options{
		HostPort: TemporalUrl,
	})
	if err != nil {
		log.Fatalln("无法连接 Temporal:", err)
	}
	defer c.Close()

	// 创建 Worker
	w := worker.New(c, "ORDER_TASK_QUEUE", worker.Options{})

	// 注册 Workflow 和 Activity
	w.RegisterWorkflow(workflows.OrderWorkflow)
	w.RegisterActivity(activities.CheckStock)
	w.RegisterActivity(activities.DeductStock)
	w.RegisterActivity(activities.ProcessPayment)
	w.RegisterActivity(activities.NotifyUser)

	// 运行 Worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("无法启动 Worker:", err)
	}
}
