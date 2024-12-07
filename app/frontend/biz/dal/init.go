package dal

import (
	"biz-demo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	//mysql.Init()
}
