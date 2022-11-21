package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 数据表对应结构体
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func Initsql() (db *gorm.DB, err error) {
	//连接数据库
	//sql: create database bubble;
	db, err = gorm.Open(mysql.Open("root:123456@tcp(124.220.204.203)/bubble"))
	return db, err
}
func main() {
	//连接数据库
	db, err := Initsql()
	if err != nil {
		fmt.Println("数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
		//延迟关闭数据库连接
		sql, sqlerr := db.DB()
		if sqlerr != nil {
			fmt.Println(sqlerr)
			return
		}
		defer sql.Close()
		//创建表
		db.AutoMigrate(&Todo{})
	}

	//启动路由
	r := gin.Default()
	//网站图标
	r.Use(favicon.New("./templates/favicon.ico"))
	//加载静态文件 模板目录替换本机目录
	r.Static("/static", "./static")
	//渲染html
	r.LoadHTMLFiles("./templates/index.html")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	//路由组
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", func(ctx *gin.Context) {
			//前端页面填写待办事项 提交请求发送到这
			//1 从请求中把数据拿出来
			var todo Todo
			ctx.BindJSON(&todo)
			//2 存入数据库
			// err = db.Create(&todo).Error
			// if err != nil {
			// }
			//3 返回响应
			if err = db.Create(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todo)
				// ctx.JSON(http.StatusOK, gin.H{
				// 	"code": 200,
				// 	"msg":  "success",
				// 	"data": todo,
				// })
			}
		})
		//查看待办事项
		v1Group.GET("/todo", func(ctx *gin.Context) {
			//查询todo表所有数据
			var todos []Todo
			if err = db.Find(&todos).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, todos)
			}
		})
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			//1 获取路径参数
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			//2 判断是否存在
			var todo Todo
			if err := db.Where("id =?", id).First(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			//3 提交数据修改
			ctx.BindJSON(&todo)
			if err = db.Save(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				ctx.JSON(http.StatusOK, todo)
			}
		})
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			if err := db.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	//端口8080
	r.Run()
}
