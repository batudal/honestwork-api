package handler

import (
	"encoding/json"
	"strconv"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/validator"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func HandleSignup(address string, signature string) string {
	salt_controller := controller.NewSaltController(address)
	salt, err := salt_controller.GetSalt()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleSignup")
		return err.Error()
	}

	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	user_controller := controller.NewUserController(address)
	existing_user, err := user_controller.GetUser()
	var user schema.User
	if err == nil {
		user = existing_user
	} else {
		conf, err := config.ParseConfig()
		if err != nil {
			return err.Error()
		}
		nft_address_hex := conf.ContractAddresses.MembershipNFT
		show_nft := validator.BoolAddr(true)
		token_id := web3.FetchUserNFT(address)
		user.ShowNFT = show_nft
		user.NFTId = strconv.Itoa(token_id)
		user.NFTAddress = nft_address_hex
	}
	user.Salt = salt

	err = user_controller.SetUser(&user)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleSignup")
		return err.Error()
	}
	return "success"
}

func HandleGetUser(address string) schema.User {
	user_controller := controller.NewUserController(address)
	user, err := user_controller.GetUser()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleGetUser")
		return schema.User{}
	}
	return user
}

func HandleUserUpdate(address string, signature string, body []byte) string {
	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	var new_user schema.User
	err := json.Unmarshal(body, &new_user)
	if err != nil {
		return err.Error()
	}

	val := validator.ValidateUserInput(&new_user, address)
	if !val {
		return err.Error()
	}

	user_controller := controller.NewUserController(address)
	existing_user, err := user_controller.GetUser()
	if err != nil {
		return err.Error()
	}

	// todo: to proper checks to see what user has mounted on profile
	new_user.Salt = existing_user.Salt
	if new_user.ImageUrl == "" {
		new_user.ImageUrl = existing_user.ImageUrl
	}

	err = user_controller.SetUser(&new_user)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleUserUpdate")
		return err.Error()
	}
	return "success"
}
