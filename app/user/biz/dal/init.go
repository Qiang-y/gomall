package dal

import (
	"biz-demo/gomall/app/user/biz/dal/mysql"
	"biz-demo/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
