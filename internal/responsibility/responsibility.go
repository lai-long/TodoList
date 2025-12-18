package responsibility

import (
	"TodoList/internal/model/dto"
	"TodoList/internal/model/entity"
	"TodoList/pkg/database"
)

func ExchangeTodo(todoInfo dto.TodoList) (todo entity.Todo) {
	todo.Context = todoInfo.Context
	todo.Status = todoInfo.Status
	todo.Title = todoInfo.Title
	todo.EndDate = todoInfo.EndDate
	return todo
}
func SaveTodo(todo entity.Todo) error {
	return database.DB.Create(&todo).Error
}
