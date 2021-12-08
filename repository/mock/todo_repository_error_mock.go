package mock

import (
	"errors"
	"todoProject/entities"
)

type TodoRepositoryError struct{}

func (repo *TodoRepositoryError) GetAllByUserId(userID int64) ([]entities.Todo, error) {
	return nil, errors.New("ERROR")
}

func (repo *TodoRepositoryError) GetByTodoIdAndUserId(todoId int64, userID int64) (entities.Todo, error) {
	return entities.Todo{}, errors.New("ERROR")
}

func (repo *TodoRepositoryError) DeleteTodo(id int64) (int64, error) {
	panic("implement me")
}

func (repo *TodoRepositoryError) Init() error {
	return nil
}

func (repo *TodoRepositoryError) CloseConnection() error {
	return nil
}

func (repo *TodoRepositoryError) InsertTodo(todo *entities.Todo) (entities.Todo, error) {
	return entities.Todo{}, errors.New("ERROR")
}

func (repo *TodoRepositoryError) UpdateTodo(todo *entities.Todo) (entities.Todo, error) {
	return entities.Todo{}, errors.New("ERROR")
}

func (repo *TodoRepositoryError) DeleteTodoById(id int64) (int64, error) {
	return 0, errors.New("ERROR")
}
