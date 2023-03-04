package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/schema"
)

type UserController struct {
	Address string
}

func NewUserController(address string) *UserController {
	return &UserController{
		Address: address,
	}
}

func (u *UserController) GetUser() (schema.User, error) {
	var user schema.User
	data, err := repository.JSONRead("user:" + u.Address)
	if err != nil {
		return schema.User{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		return schema.User{}, err
	}
	return user, nil
}

func (u *UserController) AddUser(user schema.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = repository.JSONWrite("user:"+u.Address, data)
	if err != nil {
		return err
	}
	return nil
}
