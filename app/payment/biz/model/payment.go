package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`        // 用户ID
	OrderId       string    `json:"order_id"`       // 订单ID
	TransactionId string    `json:"transaction_id"` // 交易ID
	Amount        float32   `json:"amount"`         // 支付总额
	PayAt         time.Time `json:"pay_at"`         // 支付时间
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context, payment *PaymentLog) error {
	err := db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
	return err
}
