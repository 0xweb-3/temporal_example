package workflows

import (
	"time"

	"github.com/0xweb-3/temporal_example/internal/activities"
	"go.temporal.io/sdk/workflow"
)

// 订单处理工作流
func OrderWorkflow(ctx workflow.Context, orderID string) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// 1. 验证库存
	var stockOK bool
	err := workflow.ExecuteActivity(ctx, activities.CheckStock, orderID).Get(ctx, &stockOK)
	if err != nil || !stockOK {
		return err
	}

	// 2. 扣减库存
	err = workflow.ExecuteActivity(ctx, activities.DeductStock, orderID).Get(ctx, nil)
	if err != nil {
		return err
	}

	// 3. 处理支付
	var paymentSuccess bool
	err = workflow.ExecuteActivity(ctx, activities.ProcessPayment, orderID).Get(ctx, &paymentSuccess)
	if err != nil || !paymentSuccess {
		return err
	}

	// 4. 通知用户
	err = workflow.ExecuteActivity(ctx, activities.NotifyUser, orderID).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
