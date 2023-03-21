package handler

import (
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
)

func HandleGetRating(address string) string {
	rating_controller := controller.NewRatingController(address)
	rating, err := rating_controller.GetRating()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleGetRating")
		return "0"
	}
	return rating
}
