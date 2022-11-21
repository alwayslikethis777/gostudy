package routers

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func SetupRouter() *gin.Engine {
	//启动路由
	r := gin.Default()
	//网站图标
	r.Use(favicon.New("./templates/favicon.ico"))
	//加载静态文件 模板目录替换本机目录
	r.Static("/static", "./static")
	//渲染html
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/", controller.IndexHandler)
	//路由组
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.Createtodo)
		//查看待办事项
		v1Group.GET("/todo", controller.Seletetodo)
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.Updatetodo)
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.Deletetodo)
	}
	return r
}
