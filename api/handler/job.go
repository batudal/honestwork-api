package handler

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/validator"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func HandleGetJob(address string, slot string) schema.Job {
	s, err := strconv.Atoi(slot)
	if err != nil {
		return schema.Job{}
	}
	job_controller := controller.NewJobController(address, s)
	job, err := job_controller.GetJob()
	if err != nil {
		return schema.Job{}
	}
	return job
}

func HandleGetJobs(address string) []schema.Job {
	job_index_controller := controller.NewJobIndexer("job_index")
	jobs, err := job_index_controller.GetJobs(address)
	if err != nil {
		return []schema.Job{}
	}
	return jobs
}

func HandleGetAllJobs(sort_field string, ascending bool) []schema.Job {
	job_index_controller := controller.NewJobIndexer("job_index")
	jobs, err := job_index_controller.GetAllJobs()
	if err != nil {
		return []schema.Job{}
	}
	return jobs
}

func HandleGetJobsLimit(offset int, size int) []schema.Job {
	job_index_controller := controller.NewJobIndexer("job_index")
	jobs, err := job_index_controller.GetAllJobsLimit(offset, size)
	if err != nil {
		return []schema.Job{}
	}
	return jobs
}

func HandleGetJobsTotal() int {
	return len(HandleGetAllJobs("created_at", false))
}

func HandleGetJobsFeed() []schema.Job {
	job_indexer := controller.NewJobIndexer("job_index")
	sticky_jobs, _ := job_indexer.GetAllJobsFilter("sticky_duration", 7)
	regular_jobs, _ := job_indexer.GetAllJobsFilter("sticky_duration", 1)

	var jobs []schema.Job
	for _, job := range sticky_jobs {
		jobs = append(jobs, job)
	}
	for _, job := range regular_jobs {
		jobs = append(jobs, job)
	}
	return jobs
}

func HandleAddJob() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var job schema.Job
		err := json.Unmarshal(c.Body(), &job)
		if err != nil {
			return err
		}

		transaction_controller := controller.NewTransactionController(job.TxHash)
		_, err = transaction_controller.GetTransaction()
		if err == nil {
			return fiber.NewError(fiber.StatusPaymentRequired, "Transaction consumed previously")
		}
		err = transaction_controller.AddTransaction(job.TxHash)
		if err != nil {
			return err
		}

		if err == nil {
			return err
		}

		err = validator.ValidateJobInput(&job)
		if err != nil {
			return err
		}

		job_indexer := controller.NewJobIndexer("job_index")
		existing_jobs, err := job_indexer.GetJobs(c.Params("address"))
		if err != nil {
			return err
		}
		job.Slot = len(existing_jobs)
		job.DealNetworkId = 0
		job.DealId = -1

		amount, err := web3.CalculatePayment(&job)
		if err != nil {
			return err
		}

		err = web3.CheckOutstandingPayment(c.Params("address"), job.TokenPaid, amount, job.TxHash)
		if err != nil {
			return err
		}

		job_controller := controller.NewJobController(c.Params("address"), job.Slot)
		err = job_controller.SetJob(&job)
		if err != nil {
			return err
		}
		return c.SendString("success")
	}
}

func HandleUpdateJob(address string, signature string, body []byte) string {
	var job schema.Job
	err := json.Unmarshal(body, &job)
	if err != nil {
		return err.Error()
	}

	// todo: check if a deal has started on this job
	// todo: return error if jobs doesnt exist

	s := strconv.Itoa(job.Slot)
	existing_job := HandleGetJob(address, s)
	job.Applications = existing_job.Applications
	job.CreatedAt = existing_job.CreatedAt
	job.TokenPaid = existing_job.TokenPaid
	job.TxHash = existing_job.TxHash
	if job.ImageUrl == "" {
		job.ImageUrl = existing_job.ImageUrl
	}

	err = validator.ValidateJobInput(&job)
	if err != nil {
		return err.Error()
	}

	job_controller := controller.NewJobController(address, job.Slot)
	err = job_controller.SetJob(&job)
	if err != nil {
		return err.Error()
	}
	return "success"
}

func HandleApplyJob(applicant_address string, signature string, recruiter_address string, slot string, body []byte) string {
	state := web3.FetchUserState(applicant_address)
	switch state {
	case 0:
		return "User doesn't have NFT."
	}

	var application schema.Application
	err := json.Unmarshal(body, &application)
	if err != nil {
		return err.Error()
	}
	application.Date = time.Now().Unix()

	// todo: check if a deal has started on this job
	s, err := strconv.Atoi(slot)
	if err != nil {
		return err.Error()
	}
	job_controller := controller.NewJobController(recruiter_address, s)
	existing_job, err := job_controller.GetJob()
	if err != nil {
		return err.Error()
	}

	for _, app := range existing_job.Applications {
		if app.UserAddress == applicant_address {
			return "You have already applied to this job."
		}
	}
	existing_job.Applications = append(existing_job.Applications, application)
	err = job_controller.SetJob(&existing_job)
	if err != nil {
		return err.Error()
	}

	user_controller := controller.NewUserController(applicant_address)
	existing_user, err := user_controller.GetUser()
	if err != nil {
		return err.Error()
	}
	existing_user.Applications = append(existing_user.Applications, application)
	err = user_controller.SetUser(&existing_user)
	if err != nil {
		return err.Error()
	}

	user_applications := existing_user.Applications
	recent_applications := make([]int64, 0)
	for _, app := range user_applications {
		if application.Date-app.Date < int64(time.Hour*24) {
			recent_applications = append(recent_applications, app.Date)
		}
	}

	// todo: update config and refactor this
	switch state {
	case 1:
		if len(recent_applications) > 1 {
			return "Application limit reached for tier 1"
		}
	case 2:
		if len(recent_applications) > 2 {
			return "Application limit reached for tier 2"
		}
	case 3:
		if len(recent_applications) > 4 {
			return "Application limit reached for tier 3"
		}
	}
	return "success"
}
