package router

import (
	"TodoList/internal/controlller"
	"TodoList/internal/middleware"
	"TodoList/internal/service"

	"github.com/gin-gonic/gin"

	_ "TodoList/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetRouters(r *gin.Engine) {
	//添加swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userGroup := r.Group("/user")
	{
		//注册1
		userGroup.POST("/register", controlller.Register)
		//登录2
		userGroup.POST("/login", controlller.Login)
	}
	memoGroup := r.Group("/memo", middleware.AuthConfirm)
	{
		//增加一条待办事项3
		memoGroup.POST("/wait/add", service.CreateNewList)
		//将一条待办设置为已完成4
		memoGroup.PUT("/wait/:id", service.OneUpdateToFinished)
		//将所有待办设置为已完成5
		memoGroup.PUT("/wait/all", service.AllUpdateToFinished)
		//将一条已完成设置为待办6
		memoGroup.PUT("/finished/:id", service.OneUpdateToWait)
		//将所有已完成设置为代办7
		memoGroup.PUT("/finished/all", service.AllUpdateToWait)
		//查已完成8
		memoGroup.GET("/search/finished/all", service.ShowFinishedList)
		//查未完成9
		memoGroup.GET("/search/wait/all", service.ShowWaitList)
		//查所有10
		memoGroup.GET("/all", service.ShowAllList)
		//根据关键词查＋?keyword=11
		memoGroup.GET("/search", service.ShowByKeyword)
		//删一条已完成12
		memoGroup.DELETE("/drop/finished/:id", service.DropOne)
		//删所有已完成、13
		memoGroup.DELETE("/drop/finished/all", service.DropAllFinished)
		//删一条代办14
		memoGroup.DELETE("/drop/wait/:id", service.DropOne)
		//删所有待办15
		memoGroup.DELETE("/drop/wait/all", service.DropAllWait)
		//删所有16
		memoGroup.DELETE("/drop/all", service.DropAllList)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 404,
			"msg":  "not found",
		})
	})
}
