package database

import (
	"TodoList/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() (err error) {
	DB, err = gorm.Open("mysql", config.Dsn)
	if err != nil {
		panic(err)
	}
	err = DB.DB().Ping()
	return
}
