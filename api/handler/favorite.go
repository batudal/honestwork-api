package handler

import (
	"encoding/json"
	"strconv"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetFavorites(address string) []*schema.Favorite {
	favorite_controller := controller.NewFavoriteController(address)
	favorite, err := favorite_controller.GetFavorites()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleGetFavorites")
		return nil
	}
	return favorite
}

func HandleAddFavorite(address string, signature string, body []byte) string {
	var favorite_input schema.FavoriteInput
	err := json.Unmarshal(body, &favorite_input)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleAddFavorite")
		return err.Error()
	}

	skill := HandleGetSkill(favorite_input.Address, strconv.Itoa(favorite_input.Slot))
	skill_user_controller := controller.NewUserController(skill.UserAddress)
	skill_user, err := skill_user_controller.GetUser()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleAddFavorite")
		return err.Error()
	}
	favorite := schema.Favorite{
		Input:    &favorite_input,
		Username: skill_user.Username,
		Title:    skill.Title,
		ImageUrl: skill.ImageUrls[0],
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.GetUser()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleAddFavorite")
		return err.Error()
	}
	for _, app := range user.Favorites {
		if app.Input.Address == favorite.Input.Address && app.Input.Slot == favorite.Input.Slot {
			return "You have already added this skill to favorites."
		}
	}
	user.Favorites = append(user.Favorites, &favorite)

	err = user_controller.SetUser(&user)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func HandleRemoveFavorite(address string, signature string, body []byte) string {
	var favorite_input schema.FavoriteInput
	err := json.Unmarshal(body, &favorite_input)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleRemoveFavorite")
		return err.Error()
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.GetUser()
	if err != nil {
		return "User not found."
	}
	for i, app := range user.Favorites {
		if app.Input.Address == favorite_input.Address && app.Input.Slot == favorite_input.Slot {
			user.Favorites = append(user.Favorites[:i], user.Favorites[i+1:]...)
		}
	}

	err = user_controller.SetUser(&user)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleRemoveFavorite")
		return err.Error()
	}
	return "success"
}
