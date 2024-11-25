package dal

import (
	"biz-demo/gomall/demo/demo_thrift/biz/dal/mysql"
	"biz-demo/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
