package routers

import (
	"github.com/gin-gonic/gin"
	"wanben/gin+gorm/bubble/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// v1
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateATodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 查看某一个待办事项
		v1Group.GET("/todo/:id", controller.GetATodo)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
