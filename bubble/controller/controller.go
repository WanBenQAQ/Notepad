package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"wanben/gin+gorm/bubble/models"
)

func CreateATodo(c *gin.Context) {
	// 前端页面填写待办事项，点击提交 会发请求到这里
	// 1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		// 3.返回响应
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 200,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetTodoList(c *gin.Context) {
	// 查询todo这个表中所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
    todo, err := models.GetATodo(id)
    if err != nil {
    	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
	// 后端直接修改
	//if err = DB.Model(&Todo{}).Where("id=?", id).Update("status", true).Error; err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": "删除失败"})
	//} else {
	//	c.JSON(http.StatusOK, gin.H{"success": http.StatusOK})
	//}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
