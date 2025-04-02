package activities

import (
	"context"
	"fmt"
)

// 处理支付
func ProcessPayment(ctx context.Context, orderID string) (bool, error) {
	fmt.Println("处理支付：", orderID)
	return true, nil
}
