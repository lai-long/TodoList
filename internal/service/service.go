package service

import (
	"TodoList/internal/dao"
	"TodoList/internal/middleware"
	"TodoList/internal/model/dto"
	"TodoList/internal/model/entity"
	"TodoList/pkg/database"
	"net/http"
	"strconv"
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
	todo := dao.ExchangeTodo(todoInfo)
	username := middleware.GetUsername(c)
	todo.UserName = username
	if err = dao.SaveTodo(todo); err != nil {
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
	offset := GetOffset(c)
	_, pagesize := GetPage(c)
	var todo []entity.Todo
	var todoList []dto.TodoList
	username := middleware.GetUsername(c)
	var count int64
	database.DB.Where("user_name = ? ", username).Count(&count)
	err = database.DB.Offset(offset).Limit(pagesize).Where("user_name = ? ", username).Find(&todo).Error
	todoList = dao.ExchangeTodoInfos(todo)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "success",
			"data":  todoList,
			"total": count,
			"page":  1 + offset,
		})
	}
}
func ShowFinishedList(c *gin.Context) {
	offset := GetOffset(c)
	_, pagesize := GetPage(c)
	var todo []entity.Todo
	var todoList []dto.TodoList
	username := middleware.GetUsername(c)
	err = database.DB.Offset(offset).Limit(pagesize).Where("status=? AND user_name=?", "已完成", username).Find(&todo).Error
	todoList = dao.ExchangeTodoInfos(todo)
	var count int64
	database.DB.Where("user_name = ? ", username).Count(&count)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "success",
			"data":  todoList,
			"total": count,
			"page":  1 + offset,
		})
	}
}
func ShowWaitList(c *gin.Context) {
	offset := GetOffset(c)
	_, pagesize := GetPage(c)
	var todo []entity.Todo
	var todoList []dto.TodoList
	username := middleware.GetUsername(c)
	err = database.DB.Offset(offset).Limit(pagesize).Where("status=? And user_name=?", "未完成", username).Find(&todo).Error
	todoList = dao.ExchangeTodoInfos(todo)
	var count int64
	database.DB.Where("user_name = ? ", username).Count(&count)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
			"error":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "success",
			"data":  todoList,
			"total": count,
			"page":  1 + offset,
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
	database.DB.Model(&todo).Where("id=? AND user_name=?", id, username).Updates(todo)
	todoInfo = dao.ExchangeTodoInfo(todo)
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
	database.DB.Model(&todo).Where("id=? AND user_name=?", id, username).Updates(todo)
	todoInfo = dao.ExchangeTodoInfo(todo)
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
	database.DB.Where("status =? AND user_name = ?", "未完成", username).Find(&todo)
	length := len(todo)
	todoInfo = make([]dto.TodoList, length)
	for i := 0; i < length; i++ {
		todo[i].Model.UpdatedAt = time.Now()
		todo[i].Status = "已完成"
		database.DB.Model(&todo).Where("id=?", todo[i].Model.ID).Updates(todo[i])
		todoInfo[i] = dao.ExchangeTodoInfo(todo[i])
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
	database.DB.Where("status =? AND user_name=?", "已完成", username).Find(&todo)
	length := len(todo)
	todoInfo = make([]dto.TodoList, length)
	for i := 0; i < length; i++ {
		todo[i].Model.UpdatedAt = time.Now()
		todo[i].Status = "未完成"
		database.DB.Model(&todo).Where("id=?", todo[i].Model.ID).Updates(todo[i])
		todoInfo[i] = dao.ExchangeTodoInfo(todo[i])
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": todoInfo,
	})
}
func ShowByKeyword(c *gin.Context) {
	offset := GetOffset(c)
	_, pagesize := GetPage(c)
	username := middleware.GetUsername(c)
	word := c.Query("keyword")
	var todo []entity.Todo
	database.DB.Offset(offset).Limit(pagesize).Where("title LIKE? AND user_name=?", "%"+word+"%", username).Or("context LIKE? AND user_name=?", "%"+word+"%", username).Find(&todo)
	var count int64
	database.DB.Where("title LIKE? AND user_name=?", "%"+word+"%", username).Or("context LIKE? AND user_name=?", "%"+word+"%", username).Count(&count)
	var todoList []dto.TodoList
	length := len(todo)
	todoList = make([]dto.TodoList, length)
	for i := 0; i < length; i++ {
		todoList[i] = dao.ExchangeTodoInfo(todo[i])
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "success",
		"data":  todoList,
		"total": count,
		"page":  1 + offset,
	})
}
func DropOne(c *gin.Context) {
	username := middleware.GetUsername(c)
	id := c.Params.ByName("id")
	if err = database.DB.Where("id=? And user_name=?", id, username).Delete(entity.Todo{}).Error; err != nil {
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
	if err = database.DB.Where("status = ? AND user_name=?", "已完成", username).Delete(entity.Todo{}).Error; err != nil {
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
	if err = database.DB.Where("status = ? AND user_name=?", "未完成", username).Delete(entity.Todo{}).Error; err != nil {
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
func GetPage(c *gin.Context) (page int, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	return page, pageSize
}
func GetOffset(c *gin.Context) int {
	page, pageSize := GetPage(c)
	offset := (page - 1) * pageSize
	return offset
}
