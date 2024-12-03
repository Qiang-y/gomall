package dal

import (
	"biz-demo/gomall/app/checkout/biz/dal/mysql"
	"biz-demo/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
