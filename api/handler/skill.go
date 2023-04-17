package handler

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/validator"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func HandleGetSkill() fiber.Handler {
	return func(c *fiber.Ctx) error {
		s, err := strconv.Atoi(c.Params("slot"))
		if err != nil {
			return c.JSON(schema.Skill{})
		}
		skill_controller := controller.NewSkillController(c.Params("address"), s)
		skill, err := skill_controller.GetSkill()
		if err != nil {
			return c.JSON(schema.Skill{})
		}
		return c.JSON(skill)
	}
}

func HandleGetSkills() fiber.Handler {
	return func(c *fiber.Ctx) error {
		skill_indexer := controller.NewSkillIndexer("skill_index")
		skills, err := skill_indexer.GetSkills(c.Params("address"))
		if err != nil {
			return c.JSON([]schema.Skill{})
		}
		return c.JSON(skills)
	}
}

func HandleGetPublishedSkills() fiber.Handler {
	return func(c *fiber.Ctx) error {
		skill_indexer := controller.NewSkillIndexer("skill_index")
		skills, err := skill_indexer.GetPublishedSkills(c.Params("address"))
		if err != nil {
			return c.JSON([]schema.Skill{})
		}
		return c.JSON(skills)
	}
}

func HandleGetAllSkills() fiber.Handler {
	return func(c *fiber.Ctx) error {
		skill_indexer := controller.NewSkillIndexer("skill_index")
		skills, err := skill_indexer.GetAllSkills()
		if err != nil {
			return c.JSON([]schema.Skill{})
		}
		return c.JSON(skills)
	}
}

func HandleGetSkillsLimit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		offset, _ := strconv.Atoi(c.Params("offset"))
		size, _ := strconv.Atoi(c.Params("size"))
		skill_indexer := controller.NewSkillIndexer("skill_index")
		skills, err := skill_indexer.GetAllSkillsLimit(offset, size)
		if err != nil {
			return err
		}
		return c.JSON(skills)
	}
}

func HandleGetSkillsTotal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		skill_indexer := controller.NewSkillIndexer("skill_index")
		skills, err := skill_indexer.GetAllSkills()
		if err != nil {
			return fiber.NewError(500, err.Error())
		}
		return c.JSON(len(skills))
	}
}

func HandleAddSkill() fiber.Handler {
	return func(c *fiber.Ctx) error {
		state := web3.FetchUserState(c.Params("address"))
		conf, err := config.ParseConfig()
		if err != nil {
			return err
		}
		var max_allowed int
		switch state {
		case 0:
			return fiber.NewError(500, "User doesn't have NFT.")
		case 1:
			max_allowed = conf.Settings.Skills.Tier_1
		case 2:
			max_allowed = conf.Settings.Skills.Tier_2
		case 3:
			max_allowed = conf.Settings.Skills.Tier_3
		}

		skill_indexer := controller.NewSkillIndexer("skill_index")
		skills, err := skill_indexer.GetSkills(c.Params("address"))
		if len(skills) == max_allowed {
			return fiber.ErrNotAcceptable
		}

		var skill schema.Skill
		err = json.Unmarshal(c.Body(), &skill)
		if err != nil {
			return err
		}

		skill.Slot = len(skills)
		skill.CreatedAt = time.Now().Unix()

		err = validator.ValidateSkillInput(&skill)
		if err != nil {
			return err
		}

		skill_controller := controller.NewSkillController(c.Params("address"), skill.Slot)
		err = skill_controller.SetSkill(&skill)
		if err != nil {
			return err
		}
		return c.SendString("success")
	}
}

func HandleUpdateSkill() fiber.Handler {
	return func(c *fiber.Ctx) error {
		s, err := strconv.Atoi(c.Params("slot"))
		if err != nil {
			return c.JSON(schema.Skill{})
		}
		skill_controller := controller.NewSkillController(c.Params("address"), s)
		existing_skill, err := skill_controller.GetSkill()
		state := web3.FetchUserState(c.Params("address"))
		conf, err := config.ParseConfig()
		if err != nil {
			return err
		}
		var max_allowed int
		switch state {
		case 0:
			return fiber.NewError(500, "User doesn't have NFT.")
		case 1:
			max_allowed = conf.Settings.Skills.Tier_1
		case 2:
			max_allowed = conf.Settings.Skills.Tier_2
		case 3:
			max_allowed = conf.Settings.Skills.Tier_3
		}
		if s > max_allowed-1 {
			return fiber.NewError(500, "User doesn't have that many skill slots.")
		}

		var new_skill schema.Skill
		err = json.Unmarshal(c.Body(), &new_skill)
		if err != nil {
			return err
		}

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
		new_skill.Slot = existing_skill.Slot

		err = validator.ValidateSkillInput(&new_skill)
		if err != nil {
			return err
		}

		err = skill_controller.SetSkill(&new_skill)
		if err != nil {
			return err
		}
		return c.SendString("success")
	}
}
