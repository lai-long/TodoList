package router

import (
	"TodoList/internal/controlller"
	"TodoList/internal/middleware"
	"TodoList/internal/service"

	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		//注册
		userGroup.POST("/register", controlller.Register)
		//登录
		userGroup.POST("/login", controlller.Login)
	}
	memoGroup := r.Group("/memo", middleware.AuthConfirm)
	{
		//增加一条待办事项
		memoGroup.POST("/wait/add", service.CreateNewList)
		//将一条待办设置为已完成
		memoGroup.PUT("/wait/:id", service.OneUpdateToFinished)
		//将所有待办设置为已完成
		memoGroup.PUT("/wait/all", service.AllUpdateToFinished)
		//将一条已完成设置为待办
		memoGroup.PUT("/finished/:id", service.OneUpdateToWait)
		//将所有已完成设置为代办
		memoGroup.PUT("/finished/all", service.AllUpdateToWait)
		//查已完成
		memoGroup.GET("/finished/all", service.ShowFinishedList)
		//查未完成
		memoGroup.GET("/wait/all", service.ShowWaitList)
		//查所有
		memoGroup.GET("/all", service.ShowAllList)
		//根据关键词查＋?keyword=
		memoGroup.GET("/search", service.ShowByKeyword)
		//删一条已完成
		memoGroup.DELETE("/finished/:id", service.DropOne)
		//删所有已完成、
		memoGroup.DELETE("/finished/all", service.DropAllFinished)
		//删一条代办
		memoGroup.DELETE("/wait/:id", service.DropOne)
		//删所有待办
		memoGroup.DELETE("/wait/all", service.DropAllWait)
		//删所有
		memoGroup.DELETE("/all", service.DropAllWait, service.DropAllFinished)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "not found",
		})
	})
}
