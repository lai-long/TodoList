package service

import (
	"TodoList/internal/model/dto"
	"TodoList/internal/responsibility"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(400, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todo,
		})
	}
}
