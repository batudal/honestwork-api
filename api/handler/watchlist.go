package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetWatchlist() fiber.Handler {
	return func(c *fiber.Ctx) error {
		watchlist_controller := controller.NewWatchlistController(c.Params("address"))
		watchlist, err := watchlist_controller.GetWatchlist()
		if err != nil {
			return err
		}
		return c.JSON(watchlist)
	}
}

func HandleAddWatchlist() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var watchlist_input schema.WatchlistInput
		err := json.Unmarshal(c.Body(), &watchlist_input)
		if err != nil {
			return err
		}

		job_controller := controller.NewJobController(watchlist_input.Address, watchlist_input.Slot)
		job, err := job_controller.GetJob()

		watchlist := schema.Watchlist{
			Input:    &watchlist_input,
			Username: job.Username,
			Title:    job.Title,
			ImageUrl: job.ImageUrl,
		}

		user_controller := controller.NewUserController(c.Params("address"))
		user, err := user_controller.GetUser()
		if err != nil {
			return err
		}
		for _, app := range user.Watchlist {
			if app.Input.Address == watchlist.Input.Address && app.Input.Slot == watchlist.Input.Slot {
				return fiber.NewError(fiber.StatusBadRequest, "Already exists")
			}
		}
		user.Watchlist = append(user.Watchlist, &watchlist)

		err = user_controller.SetUser(&user)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}

func HandleRemoveWatchlist() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var watchlist_input schema.WatchlistInput
		err := json.Unmarshal(c.Body(), &watchlist_input)
		if err != nil {
			return err
		}

		user_controller := controller.NewUserController(c.Params("address"))
		user, err := user_controller.GetUser()
		if err != nil {
			return err
		}
		for i, app := range user.Watchlist {
			if app.Input.Address == watchlist_input.Address && app.Input.Slot == watchlist_input.Slot {
				user.Watchlist = append(user.Watchlist[:i], user.Watchlist[i+1:]...)
			}
		}

		err = user_controller.SetUser(&user)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}
