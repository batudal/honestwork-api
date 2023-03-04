package controller

import "github.com/takez0o/honestwork-api/api/repository"

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

func (u *SaltController) AddSalt(salt string) error {
	err := repository.StringWrite("salt:"+u.Address, salt)
	if err != nil {
		return err
	}
	return nil
}

func (u *SaltController) DeleteSalt() error {
	err := repository.StringDelete("salt:" + u.Address)
	if err != nil {
		return err
	}
	return nil
}
