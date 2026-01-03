package main

import (
	"TodoList/pkg/database"
	"TodoList/router"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//@title TodoList API
//@version 1.0
//@description API

// @host 127.0.0.1/:8080
func main() {
	//连接数据库
	err := database.InitMysql()
	if err != nil {
		panic(err)
	}
	defer database.DB.Close()
	r := gin.Default()
	router.SetRouters(r)
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
