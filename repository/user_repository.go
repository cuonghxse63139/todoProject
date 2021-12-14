package repository

import (
	"log"
	"todoProject/dbConnention"
	"todoProject/entities"
)

type UserRepository interface {
	Login(user entities.User) (entities.User, error)
}

type UserRepositoryStruct struct {
	dbHandler dbConnention.DBConnection
}

func NewUserRepository(db dbConnention.DBConnection) (*UserRepositoryStruct, error) {
	repo := &UserRepositoryStruct{}
	repo.dbHandler = db
	return repo, db.Open()
}

func (repo *UserRepositoryStruct) Login(user entities.User) (entities.User, error) {
	db := repo.dbHandler.GetDb()
	row := db.QueryRow("Select * from Users where username = ? and password = ? ", user.Username, user.Password)

	var username, password string
	var id int64
	var role int

	error := row.Scan(&id, &username, &password, &role)
	if error != nil {
		log.Println(error)
		return entities.User{}, error
	}

	result := entities.User{
		Id:       id,
		Username: username,
		Password: password,
		Role:     role,
	}

	if error != nil {
		log.Println(error)
		return entities.User{}, error
	}

	error = db.Close()

	if error != nil {
		log.Println(error)
		return entities.User{}, error
	}

	return result, nil
}
