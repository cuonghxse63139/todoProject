package services

import (
	"log"
	"todoProject/entities"
	"todoProject/repository"
)

type TodoService interface {
	Init() error
	GetListTodoByUserId(userId int64) ([]entities.Todo, error)
	InsertTodo(todo *entities.Todo) (entities.Todo, error)
	DeleteTodo(id int64) (int64, error)
	UpdateTodo(todo *entities.Todo) (entities.Todo, error)
}

type TodoServiceStruct struct {
	repository repository.TodoRepository
}

func (service *TodoServiceStruct) Init() error {
	tempRepo := &repository.TodoRepositoryStruct{}
	service.repository = tempRepo
	return service.repository.Init()
}

func (service *TodoServiceStruct) InitWith(repository repository.TodoRepository) {
	service.repository = repository
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
