package dal

import (
	"biz-demo/gomall/app/product/biz/dal/mysql"
	"biz-demo/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
