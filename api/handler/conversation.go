package handler

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetConversations() fiber.Handler {
	return func(c *fiber.Ctx) error {
		conversation_controller := controller.NewConversationController(c.Params("address"))
		conversations, err := conversation_controller.GetConversations()
		if err != nil {
			return nil
		}
		return c.JSON(conversations)
	}
}

func HandleAddConversation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type input_address struct {
			MatchedUser string `json:"matched_user"`
		}
		target_user := input_address{}
		err := json.Unmarshal(c.Body(), &target_user)
		if err != nil {
			return err
		}
		target_address := target_user.MatchedUser
		if c.Params("address") == target_address {
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		target_user_controller := controller.NewUserController(target_address)
		target_user_db, err := target_user_controller.GetUser()
		if err != nil {
			if !*target_user_db.DmsOpen {
				job_indexer := controller.NewJobIndexer("jobs_index")
				user_jobs, err := job_indexer.GetJobs(c.Params("address"))
				if err != nil {
					return err
				}
				target_applied := false
				for _, job := range user_jobs {
					for _, application := range job.Applications {
						if application.UserAddress == target_address {
							target_applied = true
						}
					}
				}
				if !target_applied {
					return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
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
			MatchedUser:   c.Params("address"),
			CreatedAt:     time.Now().Unix(),
			LastMessageAt: 0,
			Muted:         false,
		}

		conversation_controller := controller.NewConversationController(c.Params("address"))
		conversations, err := conversation_controller.GetConversations()
		for _, c := range conversations {
			if c.MatchedUser == conversation.MatchedUser {
				return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
			}
		}
		conversations = append(conversations, &conversation)
		target_conversation_controller := controller.NewConversationController(target_address)
		target_conversations, err := conversation_controller.GetConversations()
		target_conversations = append(target_conversations, &target_conversation)

		err = conversation_controller.SetConversation(conversations)
		if err != nil {
			return err
		}
		err = target_conversation_controller.SetConversation(target_conversations)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}
