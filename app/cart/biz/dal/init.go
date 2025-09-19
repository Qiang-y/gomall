package dal

import (
	"github.com/Qiang-y/go-shop/app/cart/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
