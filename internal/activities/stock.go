package activities

import (
	"context"
	"fmt"
	"time"
)

// 验证库存
func CheckStock(ctx context.Context, orderID string) (bool, error) {
	fmt.Println("检查库存：", orderID)
	return true, nil
}

// 扣减库存
func DeductStock(ctx context.Context, orderID string) error {
	time.Sleep(time.Second * 12)
	fmt.Println("扣减库存：", orderID)
	return nil
}
