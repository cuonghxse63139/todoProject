package services

import (
	"log"
	"todoProject/entities"
	"todoProject/repository"
)

type TodoService interface {
	GetListTodoByUserId(userId int64) ([]entities.Todo, error)
	InsertTodo(todo *entities.Todo) (entities.Todo, error)
	DeleteTodo(id int64) (int64, error)
	UpdateTodo(todo *entities.Todo) (entities.Todo, error)
}

type TodoServiceStruct struct {
	repository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) *TodoServiceStruct {
	return &TodoServiceStruct{repository: todoRepo}
}

func (service *TodoServiceStruct) GetListTodoByUserId(userId int64) ([]entities.Todo, error) {
	todoRes, error := service.repository.GetAllByUserId(userId)
	if error != nil {
		log.Println(error)
		return nil, error
	}
	return todoRes, service.repository.CloseConnection()
}

func (service *TodoServiceStruct) InsertTodo(todo *entities.Todo) (entities.Todo, error) {
	todoRes, error := service.repository.InsertTodo(todo)
	if error != nil {
		log.Println(error)
		return entities.Todo{}, error
	}
	return todoRes, service.repository.CloseConnection()
}

func (service *TodoServiceStruct) DeleteTodo(id int64) (int64, error) {
	rowEffect, error := service.repository.DeleteTodo(id)
	if error != nil {
		log.Println(error)
		return 0, error
	}
	return rowEffect, service.repository.CloseConnection()
}

func (service *TodoServiceStruct) UpdateTodo(todo *entities.Todo) (entities.Todo, error) {
	todoRes, error := service.repository.UpdateTodo(todo)
	if error != nil {
		log.Println(error)
		return entities.Todo{}, error
	}
	return todoRes, service.repository.CloseConnection()
}
