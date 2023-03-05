package controller

import (
	"github.com/takez0o/honestwork-api/utils/schema"
)

type WatchlistController struct {
	Address string
}

func NewWatchlistController(address string) *WatchlistController {
	return &WatchlistController{
		Address: address,
	}
}

func (w *WatchlistController) GetWatchlist() ([]*schema.Watchlist, error) {
	user_controller := NewUserController(w.Address)
	user, err := user_controller.GetUser()
	if err != nil {
		return []*schema.Watchlist{}, err
	}
	return user.Watchlist, nil
}
