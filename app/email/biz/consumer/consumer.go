package consumer

import "github.com/Qiang-y/go-shop/app/email/biz/consumer/email"

// Init 统一初始化所有consumer
func Init() {
	email.ConsumerInit()
}
