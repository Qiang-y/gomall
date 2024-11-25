package mysql

import (
	"biz-demo/gomall/demo/demo_proto/biz/model"
	"biz-demo/gomall/demo/demo_proto/conf"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 设置mysqlDSN
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"))

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.User{})

	// 获取并打印mysql版本号
	fmt.Printf("%#v\n", DB.Debug().Exec("select version()"))
}
