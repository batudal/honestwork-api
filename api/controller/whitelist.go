package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/api/repository"
)

type WhitelistController struct {
}

func NewWhitelistController() *WhitelistController {
	return &WhitelistController{}
}

func (w *WhitelistController) GetWhitelist() []string {
	var whitelist []string
	data, err := repository.JSONRead("whitelist")
	if err != nil {
		return []string{}
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &whitelist)

	if err != nil {
		return []string{}
	}
	return whitelist
}
