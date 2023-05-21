package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/validator"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func HandleSignup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		salt_controller := controller.NewSaltController(c.Params("address"))
		salt, err := salt_controller.GetSalt()
		if err != nil {
			return err
		}
		state := web3.FetchUserState(c.Params("address"))
		switch state {
		case 0:
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		user_controller := controller.NewUserController(c.Params("address"))
		existing_user, err := user_controller.GetUser()
		var user schema.User
		if err == nil {
			user = existing_user
		} else {
			conf, err := config.ParseConfig()
			if err != nil {
				return err
			}
			nft_address_hex := conf.ContractAddresses.MembershipNFT
			show_nft := validator.BoolAddr(true)
			token_id := web3.FetchUserNFT(c.Params("address"))
			user.ShowNFT = show_nft
			user.NFTId = strconv.Itoa(token_id)
			user.NFTAddress = nft_address_hex
		}
		user.Salt = salt
		err = user_controller.SetUser(&user)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}

func HandleGetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user_controller := controller.NewUserController(c.Params("address"))
		user, err := user_controller.GetUser()
		if err != nil {
			return c.JSON(schema.User{})
		}
		return c.JSON(user)
	}
}

func HandleUserUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		state := web3.FetchUserState(c.Params("address"))
		switch state {
		case 0:
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
		var new_user schema.User
		err := json.Unmarshal(c.Body(), &new_user)
		if err != nil {
			return err
		}
		val := validator.ValidateUserInput(&new_user, c.Params("address"))
		if !val {
			return err
		}
		user_controller := controller.NewUserController(c.Params("address"))
		existing_user, err := user_controller.GetUser()
		if err != nil {
			return err
		}
		// todo: to proper checks to see what user has mounted on profile
		new_user.Salt = existing_user.Salt
		if new_user.ImageUrl == "" {
			new_user.ImageUrl = existing_user.ImageUrl
		}
		new_user.Rating = existing_user.Rating
		new_user.Applications = existing_user.Applications
		new_user.Watchlist = existing_user.Watchlist
		new_user.Favorites = existing_user.Favorites
		err = user_controller.SetUser(&new_user)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}
