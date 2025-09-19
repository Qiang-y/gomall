package dal

import (
	"github.com/Qiang-y/go-shop/app/product/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
