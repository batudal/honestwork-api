package handler

import (
	"github.com/takez0o/honestwork-api/api/controller"
)

func HandleWhitelist() []string {
	wl_controller := controller.NewWhitelistController()
	return wl_controller.GetWhitelist()
}
