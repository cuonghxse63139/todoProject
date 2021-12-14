package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"todoProject/dbConnention"
	"todoProject/entities"
)

type TodoRepository interface {
	GetAllByUserId(userID int64) ([]entities.Todo, error)
	GetByTodoIdAndUserId(todoId int64, userID int64) (entities.Todo, error)
	UpdateTodo(todo *entities.Todo) (entities.Todo, error)
	DeleteTodo(id int64) (int64, error)
	InsertTodo(todo *entities.Todo) (entities.Todo, error)
	CloseConnection() error
}

type TodoRepositoryStruct struct {
	dbHandler dbConnention.DBConnection
}

func NewTodoRepository(db dbConnention.DBConnection) (*TodoRepositoryStruct, error) {
	repo := &TodoRepositoryStruct{}
	repo.dbHandler = db
	return repo, db.Open()
}

func (repo *TodoRepositoryStruct) CloseConnection() error {
	error := repo.dbHandler.GetDb().Close()

	if error != nil {
		log.Println(error)
		return error
	}
	return nil
}

func (repo *TodoRepositoryStruct) GetAllByUserId(userID int64) ([]entities.Todo, error) {
	db := repo.dbHandler.GetDb()
	row, error := db.Query("Select * from ToDo where user_id = ?", userID)
	if error != nil {
		log.Println(error)
		return nil, error
	}
	return mapListResultTodoList(row)
}

func (repo *TodoRepositoryStruct) GetByTodoIdAndUserId(todoId int64, userID int64) (entities.Todo, error) {
	db := repo.dbHandler.GetDb()
	row, error := db.Query("Select * from ToDo where id = ? and user_id = ?", todoId, userID)
	if error != nil {
		log.Println(error)
		todo := entities.Todo{}
		return todo, error
	}
	return mapResultTodoList(row)
}

func (repo *TodoRepositoryStruct) InsertTodo(todo *entities.Todo) (entities.Todo, error) {
	db := repo.dbHandler.GetDb()
	row, error := db.Exec("INSERT INTO Todo(title, todo_content, status, user_id) VALUES (?, ?, ?, ?) select ID = convert(bigint, SCOPE_IDENTITY())", todo.Title, todo.Content, todo.Status, todo.UserId)

	if error != nil {
		log.Println(error)
		return entities.Todo{}, error
	}

	todo.Id, error = row.LastInsertId()

	if error != nil {
		log.Println(error)
		return entities.Todo{}, error
	}

	return repo.GetByTodoIdAndUserId(todo.Id, todo.UserId)
}

func (repo *TodoRepositoryStruct) UpdateTodo(todo *entities.Todo) (entities.Todo, error) {
	db := repo.dbHandler.GetDb()
	sqlResult, error := db.Exec("UPDATE Todo SET title = ?, todo_content = ?, status = ? WHERE id = ?", todo.Title, todo.Content, todo.Status, todo.Id)

	if error != nil {
		log.Println(error)
		return entities.Todo{}, error
	}

	rowEffected, error := sqlResult.RowsAffected()

	if error != nil {
		log.Println(error)
		return entities.Todo{}, error
	} else if rowEffected != 1 {
		log.Println("row affected is " + fmt.Sprintf("%d", rowEffected))
		return entities.Todo{}, errors.New("row affected is " + fmt.Sprintf("%d", rowEffected))
	}

	return repo.GetByTodoIdAndUserId(todo.Id, todo.UserId)
}

func (repo *TodoRepositoryStruct) DeleteTodo(id int64) (int64, error) {
	db := repo.dbHandler.GetDb()
	sqlResult, error := db.Exec("DELETE FROM Todo WHERE id = ? ", id)

	if error != nil {
		log.Println(error)
		return 0, error
	}

	rowAffect, error := sqlResult.RowsAffected()

	if error != nil {
		log.Println(error)
		return 0, error
	}

	if rowAffect == 0 {
		msg := fmt.Sprintf("Delete todo failed!!!")
		log.Println(msg)
		return 0, errors.New(msg)
	}

	return rowAffect, nil
}

func mapListResultTodoList(rows *sql.Rows) ([]entities.Todo, error) {
	var todoLst []entities.Todo

	for rows.Next() {
		var id, userId int64
		var title, content string
		var status int

		error := rows.Scan(&id, &title, &content, &status, &userId)

		if error != nil {
			log.Println(error)
			return nil, error
		}

		todo := entities.Todo{
			Id:      id,
			Title:   title,
			Content: content,
			Status:  status,
			UserId:  userId,
		}
		todoLst = append(todoLst, todo)
	}
	return todoLst, nil
}

func mapResultTodoList(rows *sql.Rows) (entities.Todo, error) {
	var todo entities.Todo
	for rows.Next() {
		var id, userId int64
		var title, content string
		var status int

		error := rows.Scan(&id, &title, &content, &status, &userId)

		if error != nil {
			log.Println(error)
			return todo, error
		}

		todo = entities.Todo{
			Id:      id,
			Title:   title,
			Content: content,
			Status:  status,
			UserId:  userId,
		}
	}
	return todo, nil
}
