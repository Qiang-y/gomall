package dal

import (
	"github.com/Qiang-y/go-shop/app/order/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
