package controller

import (
	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/schema"
)

type FavoriteController struct {
	Address string
}

func NewFavoriteController(address string) *FavoriteController {
	return &FavoriteController{
		Address: address,
	}
}

func (w *FavoriteController) GetFavorites() ([]*schema.Favorite, error) {
	user_controller := NewUserController(w.Address)
	user, err := user_controller.GetUser()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "GetFavorites - GetUser")
		return []*schema.Favorite{}, err
	}
	return user.Favorites, nil
}
