package dal

import (
	"github.com/Qiang-y/go-shop/app/checkout/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
