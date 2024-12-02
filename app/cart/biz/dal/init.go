package dal

import (
	"biz-demo/gomall/app/cart/biz/dal/mysql"
	"biz-demo/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
