package mysql

import (
	"github.com/Qiang-y/go-shop/app/order/biz/model"
	"github.com/Qiang-y/go-shop/app/order/conf"
	"fmt"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"))
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3307)/user?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 集成openTelemetry
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	if os.Getenv("GO_ENV") != "online" {
		err := DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
		if err != nil {
			panic(err)
		}
	}
}
