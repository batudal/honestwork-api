package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/parser"
	"github.com/takez0o/honestwork-api/utils/schema"
	"github.com/takez0o/honestwork-api/utils/validator"
	"github.com/takez0o/honestwork-api/utils/web3"
)

func HandleGetJob() fiber.Handler {
	return func(c *fiber.Ctx) error {
		s, err := strconv.Atoi(c.Params("slot"))
		if err != nil {
			return c.JSON(schema.Job{})
		}
		job_controller := controller.NewJobController(c.Params("address"), s)
		job, err := job_controller.GetJob()
		if err != nil {
			return c.JSON(schema.Job{})
		}
		return c.JSON(job)
	}
}

func HandleGetJobs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		job_index_controller := controller.NewJobIndexer("job_index")
		jobs, err := job_index_controller.GetJobs(c.Params("address"))
		if err != nil {
			return err
		}
		return c.JSON(jobs)
	}
}

func HandleGetAllJobs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		job_index_controller := controller.NewJobIndexer("job_index")
		jobs, err := job_index_controller.GetAllJobs()
		if err != nil {
			return c.JSON([]schema.Job{})
		}
		return c.JSON(jobs)
	}
}

func HandleGetJobsLimit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		offset, err := strconv.Atoi(c.Params("offset"))
		size, err := strconv.Atoi(c.Params("size"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}
		job_index_controller := controller.NewJobIndexer("job_index")
		jobs, err := job_index_controller.GetAllJobsLimit(offset, size)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Not Found")
		}
		return c.JSON(jobs)
	}
}

func HandleGetJobsTotal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		job_index_controller := controller.NewJobIndexer("job_index")
		jobs, err := job_index_controller.GetAllJobs()
		if err != nil {
			return err
		}
		return c.JSON(len(jobs))
	}
}

func HandleGetJobsFeed() fiber.Handler {
	return func(c *fiber.Ctx) error {
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
		return c.JSON(jobs)
	}
}

func HandleAddJob() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var job schema.Job
		err := json.Unmarshal(c.Body(), &job)
		if err != nil {
			return err
		}
		fmt.Println("Job timezone:", job.Timezone)

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

		if amount.Cmp(big.NewInt(0)) != 0 {
			transaction_controller := controller.NewTransactionController(job.TxHash)
			_, err = transaction_controller.GetTransaction()
			if err == nil {
				return fiber.NewError(fiber.StatusPaymentRequired, "Transaction consumed previously")
			}
			err = transaction_controller.AddTransaction(job.TxHash)
			if err != nil {
				return err
			}
			err = web3.CheckOutstandingPayment(c.Params("address"), job.TokenPaid, amount, job.TxHash)
			if err != nil {
				return err
			}
		}

		job_controller := controller.NewJobController(c.Params("address"), job.Slot)
		err = job_controller.SetJob(&job)
		if err != nil {
			return err
		}
		guild_id := os.Getenv("DISCORD_GUILD_JOBS")
		bot_token := os.Getenv("DISCORD_BOT_TOKEN")
		var s *discordgo.Session
		s, err = discordgo.New("Bot " + bot_token)
		if err != nil {
			log.Fatalf("Invalid bot parameters(1): %v", err)
		}
		budget := strconv.Itoa(int(job.Budget))
		timezone := strconv.Itoa(int(*job.Timezone))
		s.ChannelMessageSendEmbed(guild_id, &discordgo.MessageEmbed{
			Title:       job.Title,
			URL:         "https://honestwork.app/job/" + job.UserAddress + "/" + strconv.Itoa(job.Slot),
			Color:       0xffd369,
			Description: parser.Parse(job.Description)[:200] + "...",
			Author: &discordgo.MessageEmbedAuthor{
				Name:    job.Username,
				IconURL: job.ImageUrl,
			},
			Timestamp: time.Now().Format(time.RFC3339),
			Footer: &discordgo.MessageEmbedFooter{
				Text:    "HonestWork Job Alerts",
				IconURL: "https://honestwork-userfiles.fra1.cdn.digitaloceanspaces.com/hw-icon.png",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "🤑 Budget",
					Value:  "$" + budget,
					Inline: true,
				},
				{
					Name:   "🌍 Timezone",
					Value:  "GMT " + timezone,
					Inline: true,
				},
			},
		})

		if err != nil {
			log.Fatalf("Message send err: %v", err)
		}
		return c.JSON("success")
	}
}

func HandleUpdateJob() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var job schema.Job
		err := json.Unmarshal(c.Body(), &job)
		if err != nil {
			return err
		}

		// todo: check if a deal has started on this job
		// todo: return error if jobs doesnt exist

		job_controller := controller.NewJobController(c.Params("address"), job.Slot)
		existing_job, err := job_controller.GetJob()
		job.Applications = existing_job.Applications
		job.CreatedAt = existing_job.CreatedAt
		job.TokenPaid = existing_job.TokenPaid
		job.TxHash = existing_job.TxHash
		if job.ImageUrl == "" {
			job.ImageUrl = existing_job.ImageUrl
		}

		err = validator.ValidateJobInput(&job)
		if err != nil {
			return err
		}

		err = job_controller.SetJob(&job)
		if err != nil {
			return err
		}
		return c.SendString("success")
	}
}

func HandleApplyJob() fiber.Handler {
	return func(c *fiber.Ctx) error {
		state := web3.FetchUserState(c.Params("address"))
		switch state {
		case 0:
			return fiber.NewError(fiber.StatusPaymentRequired, "User not registered")
		}

		var application schema.Application
		err := json.Unmarshal(c.Body(), &application)
		if err != nil {
			return err
		}
		application.Date = time.Now().Unix()

		// todo: check if a deal has started on this job
		s, err := strconv.Atoi(c.Params("slot"))
		if err != nil {
			return err
		}
		job_controller := controller.NewJobController(c.Params("recruiter_address"), s)
		existing_job, err := job_controller.GetJob()
		if err != nil {
			return err
		}

		for _, app := range existing_job.Applications {
			if app.UserAddress == c.Params("address") {
				return fiber.NewError(fiber.StatusNotAcceptable, "User already applied")
			}
		}
		existing_job.Applications = append(existing_job.Applications, application)
		err = job_controller.SetJob(&existing_job)
		if err != nil {
			return err
		}

		user_controller := controller.NewUserController(c.Params("address"))
		existing_user, err := user_controller.GetUser()
		if err != nil {
			return err
		}
		existing_user.Applications = append(existing_user.Applications, application)
		err = user_controller.SetUser(&existing_user)
		if err != nil {
			return err
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
				return fiber.NewError(fiber.StatusPaymentRequired, "Application limit reached for tier 1")
			}
		case 2:
			if len(recent_applications) > 2 {
				return fiber.NewError(fiber.StatusPaymentRequired, "Application limit reached for tier 2")
			}
		case 3:
			if len(recent_applications) > 4 {
				return fiber.NewError(fiber.StatusPaymentRequired, "Application limit reached for tier 3")
			}
		}
		return c.SendString("success")
	}
}
