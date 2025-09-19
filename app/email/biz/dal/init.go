package dal

import (
	"github.com/Qiang-y/go-shop/app/email/biz/dal/mysql"
	"github.com/Qiang-y/go-shop/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
