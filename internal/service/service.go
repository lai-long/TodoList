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

// CreateNewList godoc
// @Summary      创建待办事项
// @Description  创建一条新的待办事项
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body  dto.TodoList  true  "待办事项信息"
// @Success      200  {object}  map[string]interface{}  "创建成功"
// @Failure      400  {object}  map[string]interface{}  "创建失败"
// @Router       /memo/wait/add [post]
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

// ShowAllList godoc
// @Summary      查询所有待办事项
// @Description  分页查询所有待办事项（包括已完成和未完成）
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page      query     int  false  "页码"     default(1)
// @Param        pageSize  query     int  false  "每页数量" default(10)
// @Success      200  {object}  map[string]interface{}  "查询成功"
// @Failure      400  {object}  map[string]interface{}  "查询失败"
// @Router       /memo/all [get]
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

// ShowFinishedList godoc
// @Summary      查询已完成事项
// @Description  分页查询所有已完成事项
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page      query     int  false  "页码"     default(1)
// @Param        pageSize  query     int  false  "每页数量" default(10)
// @Success      200  {object}  map[string]interface{}  "查询成功"
// @Failure      400  {object}  map[string]interface{}  "查询失败"
// @Router       /memo/search/finished/all [get]
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

// ShowWaitList godoc
// @Summary      查询未完成事项
// @Description  分页查询所有未完成事项
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page      query     int  false  "页码"     default(1)
// @Param        pageSize  query     int  false  "每页数量" default(10)
// @Success      200  {object}  map[string]interface{}  "查询成功"
// @Failure      400  {object}  map[string]interface{}  "查询失败"
// @Router       /memo/search/wait/all [get]
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

// OneUpdateToFinished godoc
// @Summary      将一条待办设置为已完成
// @Description  根据ID将单条待办事项标记为已完成
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "待办事项ID"
// @Success      200  {object}  map[string]interface{}  "更新成功"
// @Failure      400  {object}  map[string]interface{}  "更新失败"
// @Router       /memo/wait/{id} [put]
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

// OneUpdateToWait godoc
// @Summary      将一条已完成设置为待办
// @Description  根据ID将单条已完成事项标记为待办
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "待办事项ID"
// @Success      200  {object}  map[string]interface{}  "更新成功"
// @Failure      400  {object}  map[string]interface{}  "更新失败"
// @Router       /memo/finished/{id} [put]
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

// AllUpdateToFinished godoc
// @Summary      将所有待办设置为已完成
// @Description  将所有未完成事项批量标记为已完成
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "更新成功"
// @Failure      400  {object}  map[string]interface{}  "更新失败"
// @Router       /memo/wait/all [put]
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

// AllUpdateToWait godoc
// @Summary      将所有已完成设置为待办
// @Description  将所有已完成事项批量标记为待办
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "更新成功"
// @Failure      400  {object}  map[string]interface{}  "更新失败"
// @Router       /memo/finished/all [put]
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

// ShowByKeyword godoc
// @Summary      关键词搜索待办事项
// @Description  根据关键词搜索待办事项（标题或内容模糊匹配）
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        keyword   query     string  true   "搜索关键词"
// @Param        page      query     int     false  "页码"     default(1)
// @Param        pageSize  query     int     false  "每页数量" default(10)
// @Success      200  {object}  map[string]interface{}  "查询成功"
// @Failure      400  {object}  map[string]interface{}  "查询失败"
// @Router       /memo/search [get]
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

// DropOne godoc
// @Summary      删除单条待办事项
// @Description  根据ID删除单条待办事项（无论是已完成还是未完成）
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  string  true  "待办事项ID"
// @Success      200  {object}  map[string]interface{}  "删除成功"
// @Failure      400  {object}  map[string]interface{}  "删除失败"
// @Router       /memo/drop/wait/{id} [delete]
// @Router       /memo/drop/finished/{id} [delete]
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

// DropAllFinished godoc
// @Summary      删除所有已完成事项
// @Description  批量删除所有已完成事项
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "删除成功"
// @Failure      400  {object}  map[string]interface{}  "删除失败"
// @Router       /memo/drop/finished/all [delete]
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

// DropAllWait godoc
// @Summary      删除所有待办事项
// @Description  批量删除所有待办（未完成）事项
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "删除成功"
// @Failure      400  {object}  map[string]interface{}  "删除失败"
// @Router       /memo/drop/wait/all [delete]
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

// DropAllList godoc
// @Summary      删除所有事项
// @Description  删除用户的所有待办事项（包括已完成和未完成）
// @Tags         待办事项模块
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  map[string]interface{}  "删除成功"
// @Failure      400  {object}  map[string]interface{}  "删除失败"
// @Router       /memo/drop/all [delete]
func DropAllList(c *gin.Context) {
	username := middleware.GetUsername(c)
	if err = database.DB.Where("user_name=?", username).Delete(entity.Todo{}).Error; err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "fail",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

// GetPage godoc
// @Summary      获取分页参数
// @Description  从查询参数中获取页码和每页数量
// @Param        page      query  int  false  "页码"     default(1)
// @Param        pageSize  query  int  false  "每页数量" default(10)
// @return       page int, pageSize int
func GetPage(c *gin.Context) (page int, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	return page, pageSize
}

// GetOffset godoc
// @Summary      计算偏移量
// @Description  根据页码和每页数量计算数据库查询偏移量
// @Param        page      query  int  false  "页码"     default(1)
// @Param        pageSize  query  int  false  "每页数量" default(10)
// @return       offset int
func GetOffset(c *gin.Context) int {
	page, pageSize := GetPage(c)
	offset := (page - 1) * pageSize
	return offset
}
