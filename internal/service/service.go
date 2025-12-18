package service

import (
	"TodoList/internal/model/dto"
	"TodoList/internal/model/entity"
	"TodoList/internal/responsibility"
	"TodoList/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var err error

func CreateNewList(c *gin.Context) {
	var todoInfo dto.TodoList
	err := c.ShouldBindJSON(&todoInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	}
	todo := responsibility.ExchangeTodo(todoInfo)
	if err = responsibility.SaveTodo(todo); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todo,
		})
	}
}
func ShowAllList(c *gin.Context) {
	var todo []entity.Todo
	var todoList []dto.TodoList
	err = database.DB.Find(&todo).Error
	todoList = responsibility.ExchangeTodoInfos(todo)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}
func ShowFinishedList(c *gin.Context) {
	var todo []entity.Todo
	var todoList []dto.TodoList
	err = database.DB.Where("status=?", "完成").Find(&todo).Error
	todoList = responsibility.ExchangeTodoInfos(todo)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}
func ShowWaitList(c *gin.Context) {
	var todo []entity.Todo
	var todoList []dto.TodoList
	err = database.DB.Where("status=?", "未完成").Find(&todo).Error
	todoList = responsibility.ExchangeTodoInfos(todo)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}
