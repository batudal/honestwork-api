package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/validator"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func HandleGetSkill(address string, slot string) schema.Skill {
	s, err := strconv.Atoi(slot)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleGetSkill - strconv.Atoi(slot)")
		return schema.Skill{}
	}
	skill_controller := controller.NewSkillController(address, s)
	skill, err := skill_controller.GetSkill()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleGetSkill - skill_controller.GetSkill()")
		return schema.Skill{}
	}
	return skill
}

func HandleGetSkills(address string) []schema.Skill {
	skill_indexer := controller.NewSkillIndexer("skill_index")
	skills, err := skill_indexer.GetSkills(address)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleGetSkills - skill_indexer.GetSkills")
		return []schema.Skill{}
	}
	return skills
}

func HandleGetAllSkills(sort_field string, ascending bool) []schema.Skill {
	skill_indexer := controller.NewSkillIndexer("skill_index")
	skills, err := skill_indexer.GetAllSkills()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleGetAllSkills - skill_indexer.GetAllSkills")
		return []schema.Skill{}
	}
	return skills
}

func HandleGetSkillsLimit(offset int, size int) []schema.Skill {
	skill_indexer := controller.NewSkillIndexer("skill_index")
	skills, err := skill_indexer.GetAllSkillsLimit(offset, size)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleGetSkillsLimit - skill_indexer.GetAllSkillsLimit")
		return []schema.Skill{}
	}
	return skills
}

func HandleGetSkillsTotal() int {
	skill_indexer := controller.NewSkillIndexer("skill_index")
	skills, err := skill_indexer.GetAllSkills()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleGetSkillsTotal - skill_indexer.GetAllSkills")
		return 0
	}
	return len(skills)
}

func HandleAddSkill(address string, signature string, body []byte) string {
	state := web3.FetchUserState(address)
	conf, err := config.ParseConfig()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleAddSkill - config.ParseConfig()")
		return err.Error()
	}
	var max_allowed int
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		max_allowed = conf.Settings.Skills.Tier_1
	case 2:
		max_allowed = conf.Settings.Skills.Tier_2
	case 3:
		max_allowed = conf.Settings.Skills.Tier_3
	}

	all_skills := HandleGetSkills(address)
	if len(all_skills) == max_allowed {
		return "User reached skill limit."
	}

	var skill schema.Skill
	err = json.Unmarshal(body, &skill)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleAddSkill - json.Unmarshal(body, &skill)")
		return err.Error()
	}

	skill.Slot = len(all_skills)
	skill.CreatedAt = time.Now().Unix()

	err = validator.ValidateSkillInput(&skill)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleAddSkill - validator.ValidateSkillInput(&skill)")
		return err.Error()
	}

	skill_controller := controller.NewSkillController(address, skill.Slot)
	err = skill_controller.SetSkill(&skill)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleAddSkill - skill_controller.SetSkill(&skill)")
		return err.Error()
	}
	return "success"
}

func HandleUpdateSkill(address string, signature string, slot string, body []byte) string {
	existing_skill := HandleGetSkill(address, slot)
	state := web3.FetchUserState(address)
	conf, err := config.ParseConfig()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleUpdateSkill - config.ParseConfig()")
		return err.Error()
	}
	var max_allowed int
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		max_allowed = conf.Settings.Skills.Tier_1
	case 2:
		max_allowed = conf.Settings.Skills.Tier_2
	case 3:
		max_allowed = conf.Settings.Skills.Tier_3
	}
	s, _ := strconv.Atoi(slot)
	if s > max_allowed-1 {
		return "User doesn't have that many skill slots."
	}

	var new_skill schema.Skill
	err = json.Unmarshal(body, &new_skill)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleUpdateSkill - json.Unmarshal(body, &new_skill)")
		return err.Error()
	}
	fmt.Println("Skill:", new_skill)

	for index, url := range new_skill.ImageUrls {
		if url == "" {
			if len(existing_skill.ImageUrls) > index {
				new_skill.ImageUrls[index] = existing_skill.ImageUrls[index]
			} else {
				new_skill.ImageUrls[index] = ""
			}
		}
	}

	new_skill.CreatedAt = existing_skill.CreatedAt
	new_skill.UserAddress = existing_skill.UserAddress

	err = validator.ValidateSkillInput(&new_skill)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleUpdateSkill - validator.ValidateSkillInput(&new_skill)")
		return err.Error()
	}

	skill_controller := controller.NewSkillController(address, s)
	err = skill_controller.SetSkill(&new_skill)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "HandleUpdateSkill - skill_controller.SetSkill(&new_skill)")
		return err.Error()
	}
	return "success"
}
