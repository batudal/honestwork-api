package handler

import (
	"fmt"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/crypto"
)

func HandleAddSalt(address string) string {
	salt_controller := controller.NewSaltController(address)
	salt := crypto.GenerateSalt()
	fmt.Println("My salt:", salt)
	salt_controller.AddSalt(salt)
	return salt
}
