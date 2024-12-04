package dal

import (
	"biz-demo/gomall/app/email/biz/dal/mysql"
	"biz-demo/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
