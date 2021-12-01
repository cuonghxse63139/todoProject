package dtos

import "todoProject/entities"

type LoginDTO struct {
	Token string        `json:"token"`
	User  entities.User `json:"user"`
}
