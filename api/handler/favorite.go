package handler

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetFavorites() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Address:", c.Params("address"))
		favorite_controller := controller.NewFavoriteController(c.Params("address"))
		favorite, err := favorite_controller.GetFavorites()
		if err != nil {
			return nil
		}
		log.Println("Favs:", favorite)
		return c.JSON(favorite)
	}
}

func HandleAddFavorite() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var favorite_input schema.FavoriteInput
		err := json.Unmarshal(c.Body(), &favorite_input)
		if err != nil {
			return err
		}

		skill_controller := controller.NewSkillController(favorite_input.Address, favorite_input.Slot)
		skill, err := skill_controller.GetSkill()
		skill_user_controller := controller.NewUserController(skill.UserAddress)
		skill_user, err := skill_user_controller.GetUser()
		if err != nil {
			return err
		}
		favorite := schema.Favorite{
			Input:    &favorite_input,
			Username: skill_user.Username,
			Title:    skill.Title,
			ImageUrl: skill.ImageUrls[0],
		}

		user_controller := controller.NewUserController(c.Params("address"))
		user, err := user_controller.GetUser()
		if err != nil {
			return err
		}
		for _, app := range user.Favorites {
			if app.Input.Address == favorite.Input.Address && app.Input.Slot == favorite.Input.Slot {
				return fiber.NewError(fiber.StatusBadRequest, "Already exists")
			}
		}
		user.Favorites = append(user.Favorites, &favorite)

		err = user_controller.SetUser(&user)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}

func HandleRemoveFavorite() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var favorite_input schema.FavoriteInput
		err := json.Unmarshal(c.Body(), &favorite_input)
		if err != nil {
			return err
		}

		user_controller := controller.NewUserController(c.Params("address"))
		user, err := user_controller.GetUser()
		if err != nil {
			return err
		}
		for i, app := range user.Favorites {
			if app.Input.Address == favorite_input.Address && app.Input.Slot == favorite_input.Slot {
				user.Favorites = append(user.Favorites[:i], user.Favorites[i+1:]...)
			}
		}

		err = user_controller.SetUser(&user)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}
