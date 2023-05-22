package controller

import (
	"log"

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
	log.Println("Address2:", w.Address)
	user_controller := NewUserController(w.Address)
	user, err := user_controller.GetUser()
	if err != nil {
		log.Println("Err (getFavorities):", err)
		return []*schema.Favorite{}, err
	}
	return user.Favorites, nil
}
