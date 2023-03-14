package handler

import (
	"encoding/json"
	"time"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetConversations(address string) []*schema.Conversation {
	conversation_controller := controller.NewConversationController(address)
	conversations, err := conversation_controller.GetConversations()
	if err != nil {
		return nil
	}
	return conversations
}

func HandleAddConversation(address string, signature string, body []byte) string {
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

	target_user_controller := controller.NewUserController(target_address)
	target_user_db, err := target_user_controller.GetUser()
	if err != nil {
		if !*target_user_db.DmsOpen {
			job_indexer := controller.NewJobIndexer("jobs_index")
			user_jobs, err := job_indexer.GetJobs(address)
			if err != nil {
				return err.Error()
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

	conversations := HandleGetConversations(address)
	for _, c := range conversations {
		if c.MatchedUser == conversation.MatchedUser {
			return "Conversation exists already."
		}
	}
	conversations = append(conversations, &conversation)
	target_conversations := HandleGetConversations(target_address)
	target_conversations = append(target_conversations, &target_conversation)

	conversation_controller := controller.NewConversationController(address)
	err = conversation_controller.SetConversation(conversations)
	if err != nil {
		return err.Error()
	}
	target_conversation_controller := controller.NewConversationController(target_address)
	err = target_conversation_controller.SetConversation(target_conversations)
	if err != nil {
		return err.Error()
	}
	return "success"
}
