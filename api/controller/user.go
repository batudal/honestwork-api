package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/utils/schema"
  "github.com/takez0o/honestwork-api/api/repository"
)

type UserController struct {
  Address string
}

func NewUserController(address string) *UserController {
  return &UserController{
    Address: address,
  }
}

func (u *UserController) Get() (schema.User, error) {
	var user schema.User
  data,err := repository.JSONRead("user:" + u.Address)
	if err != nil {
		return schema.User{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		return schema.User{}, err
	}
	return user, nil
}
