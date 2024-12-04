package dal

import (
	"biz-demo/gomall/app/order/biz/dal/mysql"
	"biz-demo/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
