package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initsql() (err error) {
	//连接数据库
	//sql: create database bubble;
	DB, err = gorm.Open(mysql.Open("root:123456@tcp(124.220.204.203)/bubble"))
	return err
}
func Closesql() {
	sql, _ := DB.DB()
	defer sql.Close()
}
