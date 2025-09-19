package dal

import (
	"github.com/Qiang-y/go-shop/app/user/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
