package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func getSkill(redis *redis.Client, slot string, address string) schema.Skill {
	record_id := "skill:" + address + ":" + slot
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

	// only return published skills?
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

	// todo: cleanup unmarshaling
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
	record_id := "job:" + address + ":" + slot
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
	sort_field := "created_at"
	ascending := false
	data, _, err := redis.Search(redisearch.NewQuery(address).SetSortBy(sort_field, ascending).Limit(0, 10000))
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

func getAllJobs(redis *redisearch.Client, sort_field string, ascending bool) []schema.Job {
	data, _, err := redis.Search(redisearch.NewQuery("*").SetSortBy(sort_field, ascending).Limit(0, 10000))
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

func getDeals(redis *redis.Client, recruiter string, creator string) ([]schema.Deal, error) {
	record_id := "deals:" + recruiter + ":" + creator
	var deals []schema.Deal
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		fmt.Println("Cant read from redis")
		return []schema.Deal{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &deals)
	if err != nil {
		fmt.Println("Cant unmarshal")
		return []schema.Deal{}, err
	}
	fmt.Println("Deals length:", len(deals))
	return deals, nil
}

func authorizeVerify(redis *redis.Client, address string, signature string) bool {
	record_id := "user:" + address
	var user_db schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user_db)
	if err != nil {
		fmt.Println("Error:", err)
	}

	result := crypto.VerifySignature(user_db.Salt, address, signature)
	if !result {
		return false
	}
	return true
}

func authorizeVerifyWithSalt(redis *redis.Client, address string, signature string) (string, error) {
	salt_id := "salt:" + address
	salt, err := redis.Get(redis.Context(), salt_id).Result()
	if err != nil {
		return "", err
	}
	result := crypto.VerifySignature(salt, address, signature)
	if !result {
		return "", err
	}
	_, err = redis.Del(redis.Context(), salt_id).Result()
	if err != nil {
		return "", err
	}
	return salt, nil
}

func getWatchlist(redis *redis.Client, address string) []*schema.Watchlist {
	user_controller := controller.NewUserController(address)
	user, err := user_controller.Get()
	if err != nil {
		return []*schema.Watchlist{}
	}
	return user.Watchlist
}

func getFavorites(redis *redis.Client, address string) []*schema.Favorite {
	user_controller := controller.NewUserController(address)
	user, err := user_controller.Get()
	if err != nil {
		return []*schema.Favorite{}
	}
	return user.Favorites
}

func getTags(redis *redis.Client) schema.Tags {
	var tags schema.Tags
	data, err := redis.Do(redis.Context(), "JSON.GET", "tags").Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &tags)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return tags
}

func getConversations(redis *redis.Client, address string) []*schema.Conversation {
	var conversations []*schema.Conversation
	record_id := "conversations:" + address
	data, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		return []*schema.Conversation{}
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &conversations)
	if err != nil {
		return []*schema.Conversation{}
	}
	return conversations
}

func HandleSignup(redis *redis.Client, address string, signature string) string {
	salt, err := authorizeVerifyWithSalt(redis, address, signature)
	if err != nil {
		return "Invalid signature."
	}

	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	user_controller := controller.NewUserController(address)
	existing_user, err := user_controller.Get()
	var user schema.User
	if err == nil {
		user = existing_user
	} else {
		conf, err := config.ParseConfig()
		if err != nil {
			fmt.Println("Error:", err)
		}
		nft_address_hex := conf.ContractAddresses.MembershipNFT
		show_nft := boolAddr(true)
		token_id := web3.FetchUserNFT(address)
		user.ShowNFT = show_nft
		user.NFTId = strconv.Itoa(token_id)
		user.NFTAddress = nft_address_hex
	}
	user.Salt = salt
	new_data, err := json.Marshal(user)
	if err != nil {
		return err.Error()
	}

	record_id := "user:" + address
	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func HandleGetUser(redis *redis.Client, address string) schema.User {
	user_controller := controller.NewUserController(address)
	user, err := user_controller.Get()
	if err != nil {
		return schema.User{}
	}
	return user
}

func HandleUserUpdate(redis *redis.Client, address string, signature string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	state := web3.FetchUserState(address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	// new user
	var user schema.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	val := ValidateUserInput(redis, &user, address)
	if !val {
		return "Invalid input."
	}

	// current user in db
	user_controller := controller.NewUserController(address)
	existing_user, err := user_controller.Get()
	if err != nil {
		return "User not found."
	}

	// filter
	user.Salt = existing_user.Salt
	if user.ImageUrl == "" {
		user.ImageUrl = existing_user.ImageUrl
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

func HandleAddSkill(redis *redis.Client, redis_search *redisearch.Client, address string, signature string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
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

	skill.Slot = len(all_skills)
	skill.CreatedAt = time.Now().Unix()

	val := ValidateSkillInput(redis, &skill)
	if !val {
		return "Invalid input."
	}

	slot := strconv.Itoa(len(all_skills))
	record_id := "skill:" + address + ":" + slot

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

func HandleUpdateSkill(redis *redis.Client, address string, signature string, slot string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
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

	val := ValidateSkillInput(redis, &skill)
	if !val {
		return "Invalid input."
	}

	new_data, err := json.Marshal(skill)
	if err != nil {
		fmt.Println("Error:", err)
	}
	record_id := "skill:" + address + ":" + slot

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleGetJob(redis *redis.Client, address string, slot string) schema.Job {
	job := getJob(redis, address, slot)
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

func HandleAddJob(redis *redis.Client, redisearch *redisearch.Client, address string, signature string, body []byte) string {
	_, err := authorizeVerifyWithSalt(redis, address, signature)
	if err != nil {
		return "Wrong signature."
	}

	var job schema.Job
	err = json.Unmarshal(body, &job)
	if err != nil {
		return err.Error()
	}

	tx_record_id := "tx:" + job.TxHash
	_, err = redis.Get(redis.Context(), tx_record_id).Result()
	if err == nil {
		return "Transaction already consumed."
	}

	err = redis.Set(redis.Context(), tx_record_id, job.UserAddress, 0).Err()
	if err != nil {
		return "Failed to set transaction."
	}

	val := ValidateJobInput(redis, &job)
	if !val {
		return "Invalid input."
	}

	job.Applications = make([]schema.Application, 0)
	job.Slot = len(getJobs(redisearch, address))

	amount, err := web3.CalculatePayment(&job)
	if err != nil {
		return err.Error()
	}

	err = web3.CheckOutstandingPayment(address, job.TokenPaid, amount, job.TxHash)
	if err != nil {
		return err.Error()
	}

	new_data, err := json.Marshal(job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	user_jobs := getJobs(redisearch, address)
	record_id := "job:" + address + ":" + strconv.Itoa(len(user_jobs))

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		return err.Error()
	}

	return "success"
}

func HandleUpdateJob(redis *redis.Client, address string, signature string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	var job schema.Job
	err := json.Unmarshal(body, &job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// todo: check if a deal has started on this job
	// todo: return error if jobs doesnt exist

	s := strconv.Itoa(job.Slot)
	existing_job := getJob(redis, address, s)
	job.Applications = existing_job.Applications
	job.CreatedAt = existing_job.CreatedAt
	job.TokenPaid = existing_job.TokenPaid
	job.TxHash = existing_job.TxHash
	if job.ImageUrl == "" {
		job.ImageUrl = existing_job.ImageUrl
	}

	val := ValidateJobInput(redis, &job)
	if !val {
		return "Invalid input."
	}

	new_data, err := json.Marshal(job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "job:" + address + ":" + s

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleApplyJob(redis *redis.Client, applicant_address string, signature string, recruiter_address string, slot string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, applicant_address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	state := web3.FetchUserState(applicant_address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	var application schema.Application
	err := json.Unmarshal(body, &application)
	if err != nil {
		fmt.Println("Error:", err)
	}

	application.Date = time.Now().Unix()

	// check if a deal has started on this job

	existing_job := getJob(redis, recruiter_address, slot)
	for _, app := range existing_job.Applications {
		if app.UserAddress == applicant_address {
			return "You have already applied to this job."
		}
	}
	existing_job.Applications = append(existing_job.Applications, application)

	new_data, err := json.Marshal(existing_job)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "job:" + recruiter_address + ":" + slot

	user_controller := controller.NewUserController(applicant_address)
	existing_user, err := user_controller.Get()
	if err != nil {
		return "User not found."
	}
	existing_user.Applications = append(existing_user.Applications, application)
	new_applicant, err := json.Marshal(existing_user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	user_applications := existing_user.Applications
	filtered_application_dates := make([]int64, 0)
	for _, app := range user_applications {
		if application.Date-app.Date < int64(time.Hour*24) {
			filtered_application_dates = append(filtered_application_dates, app.Date)
		}
	}

	switch state {
	case 1:
		if len(filtered_application_dates) > 1 {
			return "Application limit reached for tier 1"
		}
	case 2:
		if len(filtered_application_dates) > 2 {
			return "Application limit reached for tier 2"
		}
	case 3:
		if len(filtered_application_dates) > 4 {
			return "Application limit reached for tier 3"
		}
	}

	applicant_id := "user:" + applicant_address
	redis.Do(redis.Context(), "JSON.SET", applicant_id, "$", new_applicant)
	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	return "success"
}

func HandleGetWatchlist(redis *redis.Client, address string) []*schema.Watchlist {
	watchlist := getWatchlist(redis, address)
	return watchlist
}

func HandleAddWatchlist(redis *redis.Client, address string, signature string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	var watchlist_input schema.WatchlistInput
	err := json.Unmarshal(body, &watchlist_input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	job := getJob(redis, watchlist_input.Address, strconv.Itoa(watchlist_input.Slot))

	watchlist := schema.Watchlist{
		Input:    &watchlist_input,
		Username: job.Username,
		Title:    job.Title,
		ImageUrl: job.ImageUrl,
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.Get()
	if err != nil {
		return "User not found."
	}
	for _, app := range user.Watchlist {
		if app.Input.Address == watchlist.Input.Address && app.Input.Slot == watchlist.Input.Slot {
			return "You have already added this job to watchlist."
		}
	}
	user.Watchlist = append(user.Watchlist, &watchlist)

	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "user:" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)

	return "success"
}

func HandleRemoveWatchlist(redis *redis.Client, address string, signature string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	var watchlist_input schema.WatchlistInput
	err := json.Unmarshal(body, &watchlist_input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.Get()
	if err != nil {
		return "User not found."
	}
	for i, app := range user.Watchlist {
		if app.Input.Address == watchlist_input.Address && app.Input.Slot == watchlist_input.Slot {
			user.Watchlist = append(user.Watchlist[:i], user.Watchlist[i+1:]...)
		}
	}

	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "user:" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)

	return "success"
}

func HandleGetFavorites(redis *redis.Client, address string) []*schema.Favorite {
	favorites := getFavorites(redis, address)
	return favorites
}

func HandleAddFavorite(redis *redis.Client, address string, signature string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	var favorite_input schema.FavoriteInput
	err := json.Unmarshal(body, &favorite_input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	skill := getSkill(redis, strconv.Itoa(favorite_input.Slot), favorite_input.Address)
	skill_user_controller := controller.NewUserController(skill.UserAddress)
	skill_user, err := skill_user_controller.Get()
	if err != nil {
		return "User not found."
	}
	favorite := schema.Favorite{
		Input:    &favorite_input,
		Username: skill_user.Username,
		Title:    skill.Title,
		ImageUrl: skill.ImageUrls[0],
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.Get()
	if err != nil {
		return "User not found."
	}
	for _, app := range user.Favorites {
		if app.Input.Address == favorite.Input.Address && app.Input.Slot == favorite.Input.Slot {
			return "You have already added this skill to favorites."
		}
	}
	user.Favorites = append(user.Favorites, &favorite)

	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "user:" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)

	return "success"
}

func HandleRemoveFavorite(redis *redis.Client, address string, signature string, body []byte) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	var favorite_input schema.FavoriteInput
	err := json.Unmarshal(body, &favorite_input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	user_controller := controller.NewUserController(address)
	user, err := user_controller.Get()
	if err != nil {
		return "User not found."
	}
	for i, app := range user.Favorites {
		if app.Input.Address == favorite_input.Address && app.Input.Slot == favorite_input.Slot {
			user.Favorites = append(user.Favorites[:i], user.Favorites[i+1:]...)
		}
	}

	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	record_id := "user:" + address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)

	return "success"
}

func HandleGetSalt(redis *redis.Client, address string) string {
	salt := crypto.GenerateSalt()
	salt_id := "salt:" + address
	ttl := time.Duration(24*30) * time.Hour
	err := redis.Set(redis.Context(), salt_id, salt, ttl).Err()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return salt
}

func HandleVerify(redis *redis.Client, address string, signature string) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}
	return "success"
}

func HandleGetTags(redis *redis.Client) schema.Tags {
	tags := getTags(redis)
	return tags
}

func HandleAddTag(redis *redis.Client, address string, signature string, tag string) string {
	authorizeVerifyd := authorizeVerify(redis, address, signature)
	if !authorizeVerifyd {
		return "Wrong signature."
	}

	tags := getTags(redis)

	for _, t := range tags.Tags {
		if t == tag {
			return "This tag already exists."
		}
	}
	tags.Tags = append(tags.Tags, tag)

	new_data, err := json.Marshal(tags)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", "tags", "$", new_data)

	return "success"
}

func isMember(redis *redis.Client, address string) bool {
	record_id := "user:" + address
	_, err := redis.Do(redis.Context(), "JSON.GET", record_id).Result()
	if err != nil {
		return false
	} else {
		return true
	}
}

func isApplicantOf(redis *redis.Client, recruiter string, applicant string) bool {
	return true
}

func HandleGetConversations(redis *redis.Client, address string) []*schema.Conversation {
	conversations := getConversations(redis, address)
	return conversations
}

func HandleAddConversation(redis *redis.Client, redis_search *redisearch.Client, address string, signature string, body []byte) string {
	// move to authorization module
	if isMember(redis, address) {
		authorized := authorizeVerify(redis, address, signature)
		if !authorized {
			return "Wrong signature."
		}
	} else {
		_, err := authorizeVerifyWithSalt(redis, address, signature)
		if err != nil {
			return "Wrong signature."
		}
	}

	type input_address struct {
		MatchedUser string `json:"matched_user"`
	}
	target_user := input_address{}
	err := json.Unmarshal(body, &target_user)
	if err != nil {
		return err.Error()
	}
	target_address := target_user.MatchedUser
	if address == target_address {
		return "Can't start conversation with self."
	}

	if isMember(redis, target_address) {
		user_controller := controller.NewUserController(target_address)
		target_user_db, err := user_controller.Get()
		if err != nil {
			return "Db read failed."
		}
		if !*target_user_db.DmsOpen {
			user_jobs := getJobs(redis_search, address)
			target_applied := false
			for _, job := range user_jobs {
				for _, application := range job.Applications {
					if application.UserAddress == target_address {
						target_applied = true
					}
				}
			}
			if !target_applied {
				return "User doesn't accept dms right now."
			}
		}
	}

	conversation := schema.Conversation{
		MatchedUser:   target_address,
		CreatedAt:     time.Now().Unix(),
		LastMessageAt: 0,
		Muted:         false,
	}

	target_conversation := schema.Conversation{
		MatchedUser:   address,
		CreatedAt:     time.Now().Unix(),
		LastMessageAt: 0,
		Muted:         false,
	}

	conversations := getConversations(redis, address)
	for _, c := range conversations {
		if c.MatchedUser == conversation.MatchedUser {
			return "Conversation exists already."
		}
	}
	conversations = append(conversations, &conversation)

	target_conversations := getConversations(redis, target_address)
	target_conversations = append(target_conversations, &target_conversation)

	new_data, err := json.Marshal(conversations)
	if err != nil {
		return err.Error()
	}
	target_new_data, err := json.Marshal(target_conversations)
	if err != nil {
		return err.Error()
	}

	record_id := "conversations:" + address
	target_record_id := "conversations:" + target_address

	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	redis.Do(redis.Context(), "JSON.SET", target_record_id, "$", target_new_data)

	return "success"
}

func HandleGetDeals(redis *redis.Client, recruiter string, creator string) []schema.Deal {
	deals, err := getDeals(redis, recruiter, creator)
	if err != nil {
		return []schema.Deal{}
	}
	return deals
}

func HandleAddDeal(redis *redis.Client, recruiter string, creator string, signature string, body []byte) string {
	_, err := authorizeVerifyWithSalt(redis, recruiter, signature)
	if err != nil {
		return "Invalid signature."
	}

	deals, err := getDeals(redis, recruiter, creator)
	if err != nil {
		deals = []schema.Deal{}
	}

	var deal schema.Deal
	err = json.Unmarshal(body, &deal)
	if err != nil {
		fmt.Println("Error:", err)
	}

	deal.Status = "offered"
	deals = append(deals, deal)

	new_data, err := json.Marshal(deals)
	if err != nil {
		return err.Error()
	}

	record_id := "deals:" + recruiter + ":" + creator
	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func HandleSignDeal(redis *redis.Client, recruiter string, creator string, signature string, body []byte) string {
	authorized := authorizeVerify(redis, creator, signature)
	if !authorized {
		return "Wrong signature."
	}

	deals, err := getDeals(redis, recruiter, creator)
	if err != nil {
		return err.Error()
	}

	type DealSignature struct {
		Slot      int    `json:"slot"`
		Signature string `json:"signature"`
	}

	var dealSignature DealSignature
	err = json.Unmarshal(body, &dealSignature)
	if err != nil {
		return err.Error()
	}

	if dealSignature.Slot > len(deals) {
		return "Wrong slot."
	}

	deals[dealSignature.Slot].Signature = dealSignature.Signature
	deals[dealSignature.Slot].Status = "accepted"

	new_data, err := json.Marshal(deals)
	if err != nil {
		return err.Error()
	}

	record_id := "deals:" + recruiter + ":" + creator
	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func HandleExecuteDeal(redis *redis.Client, recruiter string, creator string, signature string, body []byte) string {
	_, err := authorizeVerifyWithSalt(redis, recruiter, signature)
	if err != nil {
		return "Invalid signature."
	}

	deals, err := getDeals(redis, recruiter, creator)
	if err != nil {
		return err.Error()
	}

	type DealExecution struct {
		Slot int `json:"slot"`
	}

	var dealExecution DealExecution
	err = json.Unmarshal(body, &dealExecution)
	if err != nil {
		return err.Error()
	}

	if dealExecution.Slot > len(deals) {
		return "Wrong slot."
	}

	deals[dealExecution.Slot].Status = "executed"

	new_data, err := json.Marshal(deals)
	if err != nil {
		return err.Error()
	}

	record_id := "deals:" + recruiter + ":" + creator
	redis.Do(redis.Context(), "JSON.SET", record_id, "$", new_data)
	return "success"
}

func HandleConfig() config.Config {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return *conf
}

func getRating(redis *redis.Client, address string) (string, error) {
	record_id := "rating:" + address
	rating, err := redis.Do(redis.Context(), "GET", record_id).Result()
	if err != nil {
		return "", err
	}
	return rating.(string), nil
}

func HandleGetRating(redis *redis.Client, address string) string {
	rating, err := getRating(redis, address)
	if err != nil {
		return "Db read failed."
	}
	return rating
}
