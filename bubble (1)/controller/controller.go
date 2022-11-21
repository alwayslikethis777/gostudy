package controller

import (
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
url --->  controller  --->  logic  --->  model
请求 --->  控制器      --->  业务逻辑  --> 模型层增删改查
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func Createtodo(ctx *gin.Context) {
	//前端页面填写待办事项 提交请求发送到这
	//1 从请求中把数据拿出来
	var todo models.Todo
	ctx.BindJSON(&todo)
	//2 存入数据库
	// err = db.Create(&todo).Error
	// if err != nil {
	// }
	err := models.Createtodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
	//3 返回响应
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"msg":  "success",
	// 	"data": todo,
	// })
}
func Seletetodo(ctx *gin.Context) {
	//查询todo表所有数据
	todos, err := models.Selecttodo()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todos)
	}
}
func Updatetodo(ctx *gin.Context) {
	//1 获取路径参数
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	//2 判断是否存在
	todo, err := models.Updatecheck(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//3 提交数据修改
	ctx.BindJSON(&todo)
	err = models.Updatetodo(todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}
func Deletetodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.Deletetodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
