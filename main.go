package main

import (
	"TodoList/pkg/database"
	"TodoList/router"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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
