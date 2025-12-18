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
	r.PUT("/memo/wait")
	//将一条已完成设置为待办
	r.PUT("/memo/finished/:id")
	//将所有已完成设置为代办
	r.PUT("/memo/finished")
	//查已完成
	r.GET("/memo/finisher")
	//查未完成
	r.GET("/memo/wait")
	//查所有
	r.GET("/memo")
	//删一条已完成
	r.DELETE("/memo/finished/:id")
	//删所有已完成、
	r.DELETE("/memo/finished")
	//删一条代办
	r.DELETE("/memo/wait/:id")
	//删所有待办
	r.DELETE("/memo/wait")
	//删所有
	r.DELETE("/memo")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "not found",
		})
	})
}
