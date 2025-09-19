package dal

import (
	"github.com/Qiang-y/go-shop/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
