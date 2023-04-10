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

type JobController struct {
	Address string
	Slot    int
}

type JobIndexer struct {
	IndexName string
}

func NewJobController(address string, slot int) *JobController {
	return &JobController{
		Address: address,
		Slot:    slot,
	}
}

func NewJobIndexer(index_name string) *JobIndexer {
	return &JobIndexer{
		IndexName: index_name,
	}
}

func (j *JobController) GetJob() (schema.Job, error) {
	var job schema.Job
	data, err := repository.JSONRead("job:" + j.Address + ":" + strconv.Itoa(j.Slot))
	if err != nil {
		return schema.Job{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &job)
	if err != nil {
		return schema.Job{}, err
	}
	return job, nil
}

func (j *JobIndexer) GetJobs(address string) ([]schema.Job, error) {
	return getJobs(address, "created_at", false, 0, 10000)
}

func (j *JobIndexer) GetAllJobs() ([]schema.Job, error) {
	return getJobs("*", "created_at", false, 0, 10000)
}

func (j *JobIndexer) GetAllJobsLimit(offset int, size int) ([]schema.Job, error) {
	return getJobs("*", "created_at", false, offset, size)
}

func (j *JobIndexer) GetAllJobsFilter(filter_field string, filter_value float64) ([]schema.Job, error) {
	return getJobsFilter("created_at", false, filter_field, filter_value)
}

func (j *JobController) SetJob(job *schema.Job) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}
	// todo: update ttl -> how many days do they stay alive? 30 days?
	err = repository.JSONWrite("job:"+j.Address+":"+strconv.Itoa(j.Slot), data, 0)
	if err != nil {
		return err
	}
	return nil
}

func (j *JobController) DeleteJob() error {
	err := repository.JSONDelete("job:" + j.Address + ":" + strconv.Itoa(j.Slot))
	if err != nil {
		return err
	}
	return nil
}

func getJobs(address string, sort_field string, ascending bool, offset int, size int) ([]schema.Job, error) {
	redis := client.NewRedisSearchClient("jobIndex")
	data, _, err := redis.Search(redisearch.NewQuery(address).SetSortBy(sort_field, ascending).Limit(0, size))
	if err != nil {
		return []schema.Job{}, err
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
			return []schema.Job{}, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func getJobsFilter(sort_field string, ascending bool, filter_field string, filter_value float64) ([]schema.Job, error) {
	redis := client.NewRedisSearchClient("jobIndex")
	var f redisearch.Filter
	f.Field = filter_field
	f.Options = redisearch.NumericFilterOptions{
		Min: filter_value,
	}
	data, _, err := redis.Search(redisearch.NewQuery("*").SetSortBy(sort_field, ascending).AddFilter(f))
	if err != nil {
		panic(err)
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
			return []schema.Job{}, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
