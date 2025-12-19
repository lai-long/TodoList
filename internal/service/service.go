package service

import (
	"TodoList/internal/middleware"
	"TodoList/internal/model/dto"
	"TodoList/internal/model/entity"
	"TodoList/internal/responsibility"
	"TodoList/pkg/database"
	"net/http"
	"time"

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
	username := middleware.GetUsername(c)
	todo.UserName = username
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
	username := middleware.GetUsername(c)
	err = database.DB.Where("username = ? ", username).Find(&todo).Error
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
	username := middleware.GetUsername(c)
	err = database.DB.Where("status=? AND username=?", "已完成", username).Find(&todo).Error
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
	username := middleware.GetUsername(c)
	err = database.DB.Where("status=? And username=?", "未完成", username).Find(&todo).Error
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

func OneUpdateToFinished(c *gin.Context) {
	id := c.Params.ByName("id")
	var todoInfo dto.TodoList
	var todo entity.Todo
	todo.Model.UpdatedAt = time.Now()
	todo.Status = "已完成"
	username := middleware.GetUsername(c)
	database.DB.Model(&todo).Where("id=? AND username=?", id, username).Updates(todo)
	todoInfo = responsibility.ExchangeTodoInfo(todo)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoInfo.Status,
	})

}

func OneUpdateToWait(c *gin.Context) {
	id := c.Params.ByName("id")
	var todoInfo dto.TodoList
	var todo entity.Todo
	todo.Model.UpdatedAt = time.Now()
	todo.Status = "未完成"
	username := middleware.GetUsername(c)
	database.DB.Model(&todo).Where("id=? AND username=?", id, username).Updates(todo)
	todoInfo = responsibility.ExchangeTodoInfo(todo)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoInfo.Status,
	})
}

func AllUpdateToFinished(c *gin.Context) {
	var todo []entity.Todo
	var todoInfo []dto.TodoList
	username := middleware.GetUsername(c)
	database.DB.Where("status =? AND username = ?", "未完成", username).Find(&todo)
	length := len(todo)
	todoInfo = make([]dto.TodoList, length)
	for i := 0; i < length; i++ {
		todo[i].Model.UpdatedAt = time.Now()
		todo[i].Status = "已完成"
		database.DB.Model(&todo).Where("id=?", todo[i].Model.ID).Updates(todo[i])
		todoInfo[i] = responsibility.ExchangeTodoInfo(todo[i])
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoInfo,
	})
}

func AllUpdateToWait(c *gin.Context) {
	var todo []entity.Todo
	var todoInfo []dto.TodoList
	username := middleware.GetUsername(c)
	database.DB.Where("status =? AND username=?", "已完成", username).Find(&todo)
	length := len(todo)
	todoInfo = make([]dto.TodoList, length)
	for i := 0; i < length; i++ {
		todo[i].Model.UpdatedAt = time.Now()
		todo[i].Status = "未完成"
		database.DB.Model(&todo).Where("id=?", todo[i].Model.ID).Updates(todo[i])
		todoInfo[i] = responsibility.ExchangeTodoInfo(todo[i])
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoInfo,
	})
}

func ShowByKeyword(c *gin.Context) {
	username := middleware.GetUsername(c)
	word := c.Query("keyword")
	var todo []entity.Todo
	database.DB.Where("title LIKE? AND username=?", "%"+word+"%", username).Or("context LIKE? AND username=?", "%"+word+"%", username).Find(&todo)
	var todoList []dto.TodoList
	length := len(todo)
	todoList = make([]dto.TodoList, length)
	for i := 0; i < length; i++ {
		todoList[i] = responsibility.ExchangeTodoInfo(todo[i])
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoList,
	})
}
func DropOne(c *gin.Context) {
	username := middleware.GetUsername(c)
	id := c.Params.ByName("id")
	if err = database.DB.Where("id=? And username=?", id, username).Delete(entity.Todo{}).Error; err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func DropAllFinished(c *gin.Context) {
	username := middleware.GetUsername(c)
	if err = database.DB.Where("status = ? AND username=?", "已完成", username).DropTable(entity.Todo{}).Error; err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
func DropAllWait(c *gin.Context) {
	username := middleware.GetUsername(c)
	if err = database.DB.Where("status = ? AND username=?", "未完成", username).DropTable(entity.Todo{}).Error; err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
