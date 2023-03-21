package handler

import (
	"encoding/json"
	"strconv"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetWatchlist(address string) []*schema.Watchlist {
	watchlist_controller := controller.NewWatchlistController(address)
	watchlist, err := watchlist_controller.GetWatchlist()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
		return nil
	}
	return watchlist
}

func HandleAddWatchlist(address string, signature string, body []byte) string {
	var watchlist_input schema.WatchlistInput
	err := json.Unmarshal(body, &watchlist_input)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
		return err.Error()
	}

	job := HandleGetJob(watchlist_input.Address, strconv.Itoa(watchlist_input.Slot))

	watchlist := schema.Watchlist{
		Input:    &watchlist_input,
		Username: job.Username,
		Title:    job.Title,
		ImageUrl: job.ImageUrl,
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.GetUser()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
		return err.Error()
	}
	for _, app := range user.Watchlist {
		if app.Input.Address == watchlist.Input.Address && app.Input.Slot == watchlist.Input.Slot {
			return "You have already added this job to watchlist."
		}
	}
	user.Watchlist = append(user.Watchlist, &watchlist)

	err = user_controller.SetUser(&user)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
		return err.Error()
	}
	return "success"
}

func HandleRemoveWatchlist(address string, signature string, body []byte) string {
	var watchlist_input schema.WatchlistInput
	err := json.Unmarshal(body, &watchlist_input)
	if err != nil {
		return err.Error()
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.GetUser()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
		return err.Error()
	}
	for i, app := range user.Watchlist {
		if app.Input.Address == watchlist_input.Address && app.Input.Slot == watchlist_input.Slot {
			user.Watchlist = append(user.Watchlist[:i], user.Watchlist[i+1:]...)
		}
	}

	err = user_controller.SetUser(&user)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error())
		return err.Error()
	}
	return "success"
}
