package handler

import (
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetTags() schema.Tags {
	tag_controller := controller.NewTagController()
	tags, err := tag_controller.GetTags()
	if err != nil {
		return schema.Tags{}
	}
	return tags
}

func HandleAddTag(address string, signature string, tag string) string {
	tag_controller := controller.NewTagController()
	tags, _ := tag_controller.GetTags()

	for _, t := range tags.Tags {
		if t == tag {
			return "This tag already exists."
		}
	}
	tags.Tags = append(tags.Tags, tag)

	err := tag_controller.SetTags(&tags)
	if err != nil {
		return err.Error()
	}
	return "success"
}
