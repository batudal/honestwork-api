package handler

import "github.com/takez0o/honestwork-api/utils/config"

func HandleConfig() config.Config {
	conf, err := config.ParseConfig()
	if err != nil {
		return config.Config{}
	}
	return *conf
}
