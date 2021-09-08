package main

import (
	"wanben/gin+gorm/bubble/dao"
	"wanben/gin+gorm/bubble/models"
	"wanben/gin+gorm/bubble/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	// 2. 把模型与数据库中的表对应起来
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run(":1016")
}
