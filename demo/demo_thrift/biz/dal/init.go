package dal

import (
	"github.com/Qiang-y/go-shop/demo/demo_thrift/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
