package consumer

import "biz-demo/gomall/app/email/biz/consumer/email"

// Init 统一初始化所有consumer
func Init() {
	email.ConsumerInit()
}
