package mock

import (
	"todoProject/entities"
)

type TodoRepositorySuccess struct{}

func (repo *TodoRepositorySuccess) Init() error {
	return nil
}

func (repo *TodoRepositorySuccess) CloseConnection() error {
	return nil
}

func (repo *TodoRepositorySuccess) GetAllByUserId() ([]entities.Todo, error) {
	return []entities.Todo{}, nil
}

func (repo *TodoRepositorySuccess) InsertTodo(todo *entities.Todo) (entities.Todo, error) {
	return entities.Todo{}, nil
}

func (repo *TodoRepositorySuccess) UpdateTodo(todo *entities.Todo) (entities.Todo, error) {
	return entities.Todo{}, nil
}

func (repo *TodoRepositorySuccess) GetByTodoIdAndUserId(id int64, userId string) (entities.Todo, error) {
	return entities.Todo{}, nil
}

func (repo *TodoRepositorySuccess) DeleteTodoById(id int64) (int64, error) {
	return 0, nil
}
