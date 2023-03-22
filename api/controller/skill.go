package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/client"
	"github.com/takez0o/honestwork-api/utils/schema"
)

type SkillController struct {
	Address string
	Slot    int
}

type SkillIndexer struct {
	IndexName string
}

func NewSkillController(address string, slot int) *SkillController {
	return &SkillController{
		Address: address,
		Slot:    slot,
	}
}

func NewSkillIndexer(index_name string) *SkillIndexer {
	return &SkillIndexer{
		IndexName: index_name,
	}
}

func (s *SkillController) GetSkill() (schema.Skill, error) {
	var skill schema.Skill
	data, err := repository.JSONRead("skill:" + s.Address + ":" + strconv.Itoa(s.Slot))
	if err != nil {
		return schema.Skill{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &skill)
	if err != nil {
		return schema.Skill{}, err
	}
	return skill, nil
}

func (s *SkillController) SetSkill(skill *schema.Skill) error {
	data, err := json.Marshal(skill)
	if err != nil {
		return err
	}
	err = repository.JSONWrite("skill:"+s.Address+":"+strconv.Itoa(s.Slot), data, 0)
	if err != nil {
		return err
	}
	return nil
}

func (s *SkillIndexer) GetSkills(address string) ([]schema.Skill, error) {
	return getSkills(address, false, "created_at", true, 0, 10000)
}

func (s *SkillIndexer) GetAllSkills() ([]schema.Skill, error) {
	return getSkills("*", true, "created_at", false, 0, 10000)
}

func (s *SkillIndexer) GetAllSkillsLimit(offset int, size int) ([]schema.Skill, error) {
	return getSkills("*", true, "created_at", false, offset, size)
}

func getSkills(address string, filter bool, sort_field string, ascending bool, offset int, size int) ([]schema.Skill, error) {
	redis := client.NewRedisSearchClient("skillIndex")
	infield := "user_address"
	data, _, err := redis.Search(redisearch.NewQuery(address).SetInFields(infield).SetSortBy(sort_field, ascending).Limit(0, size))
	if err != nil {
		return []schema.Skill{}, err
	}

	var skills []schema.Skill
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			if key != sort_field {
				translationKeys = append(translationKeys, key)
			}
		}
		var skill schema.Skill
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &skill)
		if err != nil {
			return []schema.Skill{}, err
		}
		fmt.Println("Skill:", skill)
		if filter {
			if skill.Publish {
				skills = append(skills, skill)
			}
		} else {
			skills = append(skills, skill)
		}
	}
	return skills, nil
}
