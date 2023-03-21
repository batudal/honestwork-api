package handler

import (
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
)

func HandleConfig() config.Config {
	conf, err := config.ParseConfig()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleConfig")
		return config.Config{}
	}
	return *conf
}
