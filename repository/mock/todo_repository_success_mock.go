package mock

import "todoProject/entities"

type TodoRepositorySuccess struct{}

func (TodoRepositorySuccess) GetAllByUserId(userID int64) ([]entities.Todo, error) {
	return []entities.Todo{}, nil
}

func (TodoRepositorySuccess) GetByTodoIdAndUserId(todoId int64, userID int64) (entities.Todo, error) {
	return entities.Todo{}, nil
}

func (TodoRepositorySuccess) UpdateTodo(todo *entities.Todo) (entities.Todo, error) {
	return entities.Todo{}, nil
}

func (TodoRepositorySuccess) DeleteTodo(id int64) (int64, error) {
	return 1, nil
}

func (TodoRepositorySuccess) InsertTodo(todo *entities.Todo) (entities.Todo, error) {
	return entities.Todo{}, nil
}

func (TodoRepositorySuccess) CloseConnection() error {
	return nil
}
