package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/schema"
)

type TagController struct {
}

func NewTagController() *TagController {
	return &TagController{}
}

func (t *TagController) GetTags() (schema.Tags, error) {
	var tags schema.Tags
	data, err := repository.JSONRead("tags")
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "GetTags - JSONRead")
		return schema.Tags{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &tags)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "GetTags - Unmarshal")
		return schema.Tags{}, err
	}
	return tags, nil
}

func (t *TagController) SetTags(tags *schema.Tags) error {
	data, err := json.Marshal(tags)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "SetTags - Marshal")
		return err
	}
	err = repository.JSONWrite("tags", data, 0)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "SetTags - JSONWrite")
		return err
	}
	return nil
}
