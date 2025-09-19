package dal

import (
	"github.com/Qiang-y/go-shop/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	//mysql.Init()
}
