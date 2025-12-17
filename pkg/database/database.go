package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitMysql() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open db err", err)
		return
	}
	err = db.DB().Ping()
	return err
}
