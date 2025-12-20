package database

import (
	"TodoList/config"
	"TodoList/internal/model/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 初始化数据库
func InitMysql() (err error) {
	DB, err = gorm.Open("mysql", config.Dsn)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.Todo{})
	err = DB.DB().Ping()
	return
}
