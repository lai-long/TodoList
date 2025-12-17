package database

import (
	"TodoList/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitMysql() (err error) {
	db, err = gorm.Open("mysql", config.Dsn)
	if err != nil {
		fmt.Println("open db err", err)
		return
	}
	err = db.DB().Ping()

	return err
}
