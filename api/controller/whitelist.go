package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
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
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "GetWhitelist - JSONRead")
		return []string{}
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &whitelist)

	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "GetWhitelist - Unmarshal")
		return []string{}
	}
	return whitelist
}
