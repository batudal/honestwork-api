package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/schema"
)

type JobController struct {
	Address string
	Slot    int
}

func NewJobController(address string, slot int) *JobController {
	return &JobController{
		Address: address,
		Slot:    slot,
	}
}

func (u *JobController) Get() (schema.Job, error) {
	var job schema.Job
	data, err := repository.JSONRead("job:" + u.Address + ":" + strconv.Itoa(u.Slot))
	if err != nil {
		return schema.Job{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &job)
	if err != nil {
		return schema.Job{}, err
	}
	return job, nil
}
