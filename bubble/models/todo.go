package models

import "bubble/sql"

// 数据表对应结构体
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func Createtable() {
	//创建表
	sql.DB.AutoMigrate(&Todo{})
}

//增删改查都在models
func Createtodo(todo *Todo) (err error) {
	err = sql.DB.Create(&todo).Error
	return err

}
func Selecttodo() (todos []*Todo, err error) {
	if err = sql.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return
}
func Updatecheck(id string) (todo *Todo, err error) {
	if err = sql.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}
func Updatetodo(todo *Todo) (err error) {
	err = sql.DB.Save(&todo).Error
	return
}
func Deletetodo(id string) (err error) {
	err = sql.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
