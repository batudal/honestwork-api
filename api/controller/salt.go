package controller

import (
	"time"

	"github.com/takez0o/honestwork-api/api/repository"
)

type SaltController struct {
	Address string
}

func NewSaltController(address string) *SaltController {
	return &SaltController{
		Address: address,
	}
}

func (u *SaltController) GetSalt() (string, error) {
	salt, err := repository.StringRead("salt:" + u.Address)
	if err != nil {
		return "", err
	}
	return salt, nil
}

func (u *SaltController) AddSalt(salt string) (string, error) {
	salt_id := "salt:" + u.Address
	ttl := time.Duration(24*15) * time.Hour
	err := repository.StringWrite(salt_id, salt, ttl)
	if err != nil {
		return "", err
	}
	return salt, nil
}

func (u *SaltController) DeleteSalt() error {
	err := repository.StringDelete("salt:" + u.Address)
	if err != nil {
		return err
	}
	return nil
}
