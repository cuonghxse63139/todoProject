package services_test

import (
	"testing"
	"todoProject/entities"
	"todoProject/repository/mock"
	"todoProject/services"
)

func TestGetListTodoByUserIdError1(t *testing.T) {
	repo := mock.TodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.GetListTodoByUserId(0)
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestGetListTodoByUserIdError2(t *testing.T) {
	repo := mock.TodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.GetListTodoByUserId(100)
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestInsertTodoError1(t *testing.T) {
	repo := mock.TodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.InsertTodo(&entities.Todo{})
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestUpdateTodoError1(t *testing.T) {
	repo := mock.TodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.UpdateTodo(nil)
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestDeleteTodoError1(t *testing.T) {
	repo := mock.TodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.DeleteTodo(0)
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestGetListTodoByUserIdSuccess1(t *testing.T) {
	repo := mock.TodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.GetListTodoByUserId(1)
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestInsertTodoSuccess1(t *testing.T) {
	repo := mock.TodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.InsertTodo(&entities.Todo{})
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestUpdateTodoSuccess1(t *testing.T) {
	repo := mock.TodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.UpdateTodo(&entities.Todo{})
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestDeleteTodoSuccess1(t *testing.T) {
	repo := mock.TodoRepositorySuccess{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	_, error := service.DeleteTodo(1)
	if error == nil {
		t.Error("Expect error but got nil")
	}
}

func TestInit(t *testing.T) {
	repo := mock.TodoRepositoryError{}
	service := services.TodoServiceStruct{}
	service.InitWith(&repo)
	service.Init()
}
