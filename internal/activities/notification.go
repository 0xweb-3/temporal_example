package activities

import (
	"context"
	"fmt"
)

// 通知用户
func NotifyUser(ctx context.Context, orderID string) error {
	fmt.Println("通知用户：", orderID)
	return nil
}
