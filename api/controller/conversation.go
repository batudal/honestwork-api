package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/schema"
)

type ConversationController struct {
	Address string
}

func NewConversationController(address string) *ConversationController {
	return &ConversationController{
		Address: address,
	}
}

func (c *ConversationController) GetConversations() ([]*schema.Conversation, error) {
	var conversations []*schema.Conversation
	data, err := repository.JSONRead("conversations:" + c.Address)
	if err != nil {
		return []*schema.Conversation{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &conversations)
	if err != nil {
		return []*schema.Conversation{}, err
	}
	return conversations, nil
}

func (c *ConversationController) SetConversation(conversations *schema.Conversation) error {
	data, err := json.Marshal(conversations)
	if err != nil {
		return err
	}
	err = repository.JSONWrite("conversations:"+c.Address, data, 0)
	if err != nil {
		return err
	}
	return nil
}
