package main

import (
	"TodoList/pkg/database"
	"TodoList/router"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := database.InitMysql()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	router.SetRouters(r)
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
