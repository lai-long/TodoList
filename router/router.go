package router

import (
	"TodoList/internal/service"

	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	//增加一条待办事项
	r.POST("/memo/wait/add", service.CreateNewList)
	//将一条待办设置为已完成
	r.PUT("/memo/wait/:id", service.OneUpdateToFinished)
	//将所有待办设置为已完成
	r.PUT("/memo/wait/all", service.AllUpdateToFinished)
	//将一条已完成设置为待办
	r.PUT("/memo/finished/:id", service.OneUpdateToWait)
	//将所有已完成设置为代办
	r.PUT("/memo/finished/all", service.AllUpdateToWait)
	//查已完成
	r.GET("/memo/finished/all", service.ShowFinishedList)
	//查未完成
	r.GET("/memo/wait/all", service.ShowWaitList)
	//查所有
	r.GET("/memo/all", service.ShowAllList)
	//根据关键词查＋?keyword=
	r.GET("/search", service.ShowByKeyword)
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
