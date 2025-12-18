package router

import (
	"TodoList/internal/service"

	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	//增加一条待办事项
	r.POST("/memo/wait/add", service.CreateNewList)
	//将一条设置为已完成
	r.PUT("/memo/wait/:id")
	//将所有设置为已完成
	r.PUT("/memo/wait/all")
	//将一条已完成设置为待办
	r.PUT("/memo/finished/:id")
	//将所有已完成设置为代办
	r.PUT("/memo/finished/all")
	//查已完成
	r.GET("/memo/finished/all")
	//查未完成
	r.GET("/memo/wait/all")
	//查所有
	r.GET("/memo/all", service.ShowList)
	//删一条已完成
	r.DELETE("/memo/finished/:id")
	//删所有已完成、
	r.DELETE("/memo/finished/all")
	//删一条代办
	r.DELETE("/memo/wait/:id")
	//删所有待办
	r.DELETE("/memo/wait/all")
	//删所有
	r.DELETE("/memo/all")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "not found",
		})
	})
}
