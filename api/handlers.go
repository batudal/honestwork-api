package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/web3"
)

// todo: move validation to package
// todo: create data extractor (for query results)
func getUser(redis *redis.Client, address string) schema.User {
	record_id := "user:" + address
	var user schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user
}

func getSkill(redis *redis.Client, slot string, address string) schema.Skill {
	record_id := "skill:" + slot + ":" + address
	fmt.Println("record_id:", record_id)
	var skill schema.Skill
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return skill
}

func getSkills(redis *redisearch.Client, address string) []schema.Skill {
	sort_field := "created_at"
	infield := "user_address"
	data, _, err := redis.Search(redisearch.NewQuery(address).SetInFields(infield).SetSortBy(sort_field, true))

	if err != nil {
		fmt.Println("Error:", err)
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
			fmt.Println("Error:", err)
		}
		if skill.Publish {
			skills = append(skills, skill)
		}
	}
	return skills
}

func getAllSkills(redis *redisearch.Client, sort_field string, ascending bool) []schema.Skill {
	data, _, err := redis.Search(redisearch.NewQuery("*").SetSortBy(sort_field, ascending))
	if err != nil {
		fmt.Println("Error:", err)
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
			fmt.Println("Error:", err)
		}
		if skill.Publish {
			skills = append(skills, skill)
		}
	}
	return skills
}

func getSkillsLimit(redis *redisearch.Client, offset int, size int) []schema.Skill {
	sort_field := "created_at"
	data, _, err := redis.Search(redisearch.NewQuery("*").Limit(offset, size).SetSortBy(sort_field, false))
	if err != nil {
		fmt.Println("Error:", err)
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
			fmt.Println("Error:", err)
		}
		skills = append(skills, skill)
	}

	return skills
}

func getJobsLimit(redis *redisearch.Client, offset int, size int) []schema.Job {
	sort_field := "created_at"
	data, _, err := redis.Search(redisearch.NewQuery("*").Limit(offset, size).SetSortBy(sort_field, false))
	if err != nil {
		fmt.Println("Error:", err)
	}

	var jobs []schema.Job
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			if key != sort_field {
				translationKeys = append(translationKeys, key)
			}
		}
		var job schema.Job
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &job)
		if err != nil {
			fmt.Println("Error:", err)
		}
		jobs = append(jobs, job)
	}

	return jobs
}

func getTotalSkills(redis *redisearch.Client) int {
	_, total, err := redis.Search(redisearch.NewQuery("*").Limit(0, 0))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return total
}

func getTotalJobs(redis *redisearch.Client) int {
	_, total, err := redis.Search(redisearch.NewQuery("*").Limit(0, 0))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return total
}

func getJob(redis *redis.Client, address string, slot string) schema.Job {
	record_id := "job:" + slot + ":" + address
	fmt.Println("record_id:", record_id)
	var job schema.Job
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &job)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return job
}

func getJobs(redis *redisearch.Client, address string) []schema.Job {
	sort_field := "budget"
	infield := "user_address"
	data, _, err := redis.Search(redisearch.NewQuery(address).SetInFields(infield).SetSortBy(sort_field, true))

	if err != nil {
		fmt.Println("Error:", err)
	}

	var jobs []schema.Job
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			if key != sort_field {
				translationKeys = append(translationKeys, key)
			}
		}
		var job schema.Job
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &job)
		if err != nil {
			fmt.Println("Error:", err)
		}
		jobs = append(jobs, job)
	}
	return jobs
}

func getAllowedSkillAmount(tier int) int {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}
	switch tier {
	case 1:
		return conf.Settings.Skills.Tier_1
	case 2:
		return conf.Settings.Skills.Tier_2
	case 3:
		return conf.Settings.Skills.Tier_3
	default:
		return 0
	}
}

// todo: implement all validators
func validateUserInput(redis *redis.Client, user schema.User) bool {
	if ValidateUsername(user.Username) &&
		ValidateTitle(user.Title) &&
		ValidateBio(user.Bio) {
		return true
	}
	return false
}

func authorize(redis *redis.Client, address string, salt string, signature string) bool {
	result := crypto.VerifySignature(salt, address, signature)
	if result {
		return AuthorizeSignature(redis, address, salt, signature)
	}
	return false
}

func HandleSignup(redis *redis.Client, address string, salt string, signature string) string {
	result := crypto.VerifySignature(salt, address, signature)
	if !result {
		return "Wrong signature."
	}
	// new user
	user := getUser(redis, address)

	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	user.Salt = salt
	user.Signature = signature

	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "user:" + address

	fmt.Println("Updating user:", record_id)

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleGetUser(redis *redis.Client, address string) schema.User {
	user := getUser(redis, address)
	return user
}

func HandleUserUpdate(redis *redis.Client, address string, salt string, signature string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		return "User didn't bind NFT yet."
	}

	// new user
	var user schema.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if !validateUserInput(redis, user) {
		return "Invalid input."
	}

	// current user in db
	user_db := getUser(redis, address)

	// filter
	user.Salt = user_db.Salt
	user.Signature = user_db.Signature
	if user.ImageUrl == "" {
		user.ImageUrl = user_db.ImageUrl
	}

	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "user:" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleGetSkill(redis *redis.Client, address string, slot string) schema.Skill {
	skill := getSkill(redis, slot, address)
	return skill
}

func HandleGetSkills(redis *redisearch.Client, address string) []schema.Skill {
	skills := getSkills(redis, address)
	return skills
}

func HandleGetAllSkills(redis *redisearch.Client, sort_field string, ascending bool) []schema.Skill {
	skills := getAllSkills(redis, sort_field, ascending)
	return skills
}

func HandleGetSkillsLimit(redis *redisearch.Client, offset int, size int) []schema.Skill {
	data := getSkillsLimit(redis, offset, size)
	return data
}

func HandleGetSkillsTotal(redis *redisearch.Client) int {
	data := getTotalSkills(redis)
	return data
}

func HandleAddSkill(redis *redis.Client, redis_search *redisearch.Client, address string, salt string, signature string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	state := web3.FetchUserState(address)
	var max_allowed int
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		return "User didn't bind NFT yet."
	case 2:
		max_allowed = getAllowedSkillAmount(1)
	case 3:
		max_allowed = getAllowedSkillAmount(2)
	case 4:
		max_allowed = getAllowedSkillAmount(3)
	}

	all_skills := getSkills(redis_search, address)
	if len(all_skills) == max_allowed {
		return "User reached skill limit."
	}

	var skill schema.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	slot := strconv.Itoa(len(all_skills))
	record_id := "skill:" + slot + ":" + address

	fmt.Println("Record id:", record_id)
	fmt.Println("Skill:", skill)
	new_data, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleUpdateSkill(redis *redis.Client, address string, salt string, signature string, slot string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	current_skill := getSkill(redis, slot, address)
	state := web3.FetchUserState(address)
	var max_allowed int
	switch state {
	case 0:
		return "User doesn't have NFT."
	case 1:
		return "User didn't bind NFT yet."
	case 2:
		max_allowed = getAllowedSkillAmount(1)
	case 3:
		max_allowed = getAllowedSkillAmount(2)
	case 4:
		max_allowed = getAllowedSkillAmount(3)
	}

	s, _ := strconv.Atoi(slot)
	if s > max_allowed-1 {
		return "User doesn't have that many skill slots."
	}

	var skill schema.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for index, url := range skill.ImageUrls {
		if url == "" {
			if len(current_skill.ImageUrls) > index {
				skill.ImageUrls[index] = current_skill.ImageUrls[index]
			} else {
				skill.ImageUrls[index] = ""
			}
		}
	}

	skill.CreatedAt = current_skill.CreatedAt
	skill.UserAddress = current_skill.UserAddress

	new_data, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error:", err)
	}
	record_id := "skill:" + slot + ":" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func getAllJobs(redis *redisearch.Client, sort_field string, ascending bool) []schema.Job {
	data, _, err := redis.Search(redisearch.NewQuery("*").SetSortBy(sort_field, ascending))
	if err != nil {
		fmt.Println("Error:", err)
	}

	var jobs []schema.Job
	for _, d := range data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			if key != sort_field {
				translationKeys = append(translationKeys, key)
			}
		}
		var job schema.Job
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &job)
		if err != nil {
			fmt.Println("Error:", err)
		}
		jobs = append(jobs, job)
	}
	return jobs
}

func HandleGetJob(redis *redis.Client, address string, slot string) schema.Job {
	job := getJob(redis, slot, address)
	return job
}

func HandleGetJobs(redis *redisearch.Client, address string) []schema.Job {
	jobs := getJobs(redis, address)
	return jobs
}

func HandleGetAllJobs(redis *redisearch.Client, sort_field string, ascending bool) []schema.Job {
	jobs := getAllJobs(redis, sort_field, ascending)
	return jobs
}

func HandleGetJobsLimit(redis *redisearch.Client, offset int, size int) []schema.Job {
	jobs := getJobsLimit(redis, offset, size)
	return jobs
}

func HandleGetJobsTotal(redis *redisearch.Client) int {
	jobs := getTotalJobs(redis)
	return jobs
}

func HandleGetJobsFeed(redis *redisearch.Client) []schema.Job {
	sort_field := "created_at"
	filter_field := "sticky_duration"
	var f redisearch.Filter
	f.Field = filter_field
	// todo: fetch from config
	f.Options = redisearch.NumericFilterOptions{
		Min: 7,
	}
	sticky_data, _, err := redis.Search(redisearch.NewQuery("*").SetSortBy(sort_field, false).AddFilter(f))
	if err != nil {
		fmt.Println("Error:", err)
	}

	f.Options = redisearch.NumericFilterOptions{
		Max: 1,
	}
	regular_data, _, err := redis.Search(redisearch.NewQuery("*").SetSortBy(sort_field, false).AddFilter(f))
	if err != nil {
		fmt.Println("Error:", err)
	}

	var jobs []schema.Job

	var sticky_jobs []schema.Job
	for _, d := range sticky_data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			if key != sort_field {
				translationKeys = append(translationKeys, key)
			}
		}
		var job schema.Job
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &job)
		if err != nil {
			fmt.Println("Error:", err)
		}
		jobs = append(sticky_jobs, job)
	}

	var regular_jobs []schema.Job
	for _, d := range regular_data {
		translationKeys := make([]string, 0, len(d.Properties))
		for key := range d.Properties {
			if key != sort_field {
				translationKeys = append(translationKeys, key)
			}
		}
		var job schema.Job
		err = json.Unmarshal([]byte(fmt.Sprint(d.Properties[translationKeys[0]])), &job)
		if err != nil {
			fmt.Println("Error:", err)
		}
		jobs = append(regular_jobs, job)
	}

	return jobs
}

func HandleAddJob(redis *redis.Client, redisearch *redisearch.Client, address string, salt string, signature string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	var job schema.Job
	err := json.Unmarshal(body, &job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	amount, err := web3.CalculatePayment(&job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = web3.CheckOutstandingPayment(address, job.TokenPaid, amount)
	if err != nil {
		return err.Error()
	}

	new_data, err := json.Marshal(job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	user_jobs := getJobs(redisearch, address)
	record_id := "job:" + strconv.Itoa(len(user_jobs)) + ":" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleUpdateJob(redis *redis.Client, address string, salt string, signature string, slot string, body []byte) string {
	authorized := authorize(redis, address, salt, signature)
	if !authorized {
		return "Wrong signature."
	}

	var job schema.Job
	err := json.Unmarshal(body, &job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	new_data, err := json.Marshal(job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "job:" + slot + ":" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}
