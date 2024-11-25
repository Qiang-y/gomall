package main

import (
	"biz-demo/gomall/demo/demo_proto/biz/dal"
	"biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"biz-demo/gomall/demo/demo_proto/biz/model"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CRUD
	//mysql.DB.Create(&model.User{
	//	Email:    "demo@example.com",
	//	Password: "passworddemo",
	//})
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").Update("password", "2232332")
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example.com").First(&row)
	fmt.Printf("%#v", row)

	mysql.DB.Where("email = ?", "demo@example.com").Delete(&model.User{})            // 软删
	mysql.DB.Unscoped().Where("email = ?", "demo@example.com").Delete(&model.User{}) // 硬删

}
