package dal

import (
	"github.com/Qiang-y/go-shop/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
