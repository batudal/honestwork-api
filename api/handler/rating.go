package handler

import "github.com/takez0o/honestwork-api/api/controller"

func HandleGetRating(address string) string {
	rating_controller := controller.NewRatingController(address)
	rating, err := rating_controller.GetRating()
	if err != nil {
		return "0"
	}
	return rating
}
