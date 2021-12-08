package services

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
	"todoProject/config"
	"todoProject/dtos"
	"todoProject/entities"
	"todoProject/repository"
)

type UserAccountService interface {
	Init() error
	Login(user entities.User) (dtos.LoginDTO, error)
}

type UserAccountServiceStruct struct {
	repository repository.UserRepository
}

func (service *UserAccountServiceStruct) Init() error {
	tempRepo := &repository.UserRepositoryStruct{}
	service.repository = tempRepo
	return service.repository.Init()
}

func (service *UserAccountServiceStruct) InitWith(repository repository.UserRepository) {
	service.repository = repository
}

func (service *UserAccountServiceStruct) Login(user entities.User) (dtos.LoginDTO, error) {
	userRes, error := service.repository.Login(user)
	if error != nil {
		log.Println(error)
		return dtos.LoginDTO{}, error
	}
	token, error := createToken(userRes.Id, userRes.Username, userRes.Role)
	if error != nil {
		log.Println(error)
		return dtos.LoginDTO{}, error
	}
	return dtos.LoginDTO{Token: token, User: userRes}, nil
}

func createToken(id int64, loginId string, role int) (string, error) {
	var err error
	os.Setenv("ACCESS_SECRET", config.SECRET_JWT_TOKEN)
	atClaims := jwt.MapClaims{}
	atClaims[config.TOKEN_CURRENT_USER_ID] = id
	atClaims[config.TOKEN_CURRENT_USERNAME] = loginId
	atClaims[config.TOKEN_CURRENT_USER_ROLE] = role
	atClaims[config.TOKEN_EXP] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
