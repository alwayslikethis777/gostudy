package main

import (
	"bubble/models"
	"bubble/routers"
	"bubble/sql"
)

func main() {
	//连接数据库
	err := sql.Initsql()
	if err != nil {
		panic(err)
	}
	//模型绑定
	models.Createtable()
	//关闭数据库
	defer sql.Closesql()

	r := routers.SetupRouter()
	//端口80
	r.Run("80")
}
