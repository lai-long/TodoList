package dao

import (
	"TodoList/internal/model/dto"
	"TodoList/internal/model/entity"
	"TodoList/pkg/database"
)

// 将dto转化为entity
func ExchangeTodo(todoInfo dto.TodoList) (todo entity.Todo) {
	todo.Context = todoInfo.Context
	todo.Status = todoInfo.Status
	todo.Title = todoInfo.Title
	todo.EndDate = todoInfo.EndDate
	return todo
}

// 将entity转化为dto
func ExchangeTodoInfo(todo entity.Todo) (todoList dto.TodoList) {
	todoList.Title = todo.Title
	todoList.Status = todo.Status
	todoList.Context = todo.Context
	todoList.EndDate = todo.EndDate
	return todoList
}
func ExchangeTodoInfos(todo []entity.Todo) (todoList []dto.TodoList) {
	length := len(todo)
	todoList = make([]dto.TodoList, length)
	for i := 0; i < length; i++ {
		todoList[i] = ExchangeTodoInfo(todo[i])
	}
	return todoList
}

// 保存进数据库
func SaveTodo(todo entity.Todo) error {
	return database.DB.Create(&todo).Error
}
