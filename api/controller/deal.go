package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/schema"
)

type DealController struct {
	RecruiterAddress string
	CreatorAddress   string
}

type DealIndexer struct {
	IndexName string
}

func NewDealController(recruiter_address string, creator_address string) *DealController {
	return &DealController{
		RecruiterAddress: recruiter_address,
		CreatorAddress:   creator_address,
	}
}

func (s *DealController) GetDeals() ([]*schema.Deal, error) {
	var deals []*schema.Deal
	deal, err := repository.JSONRead("deals:" + s.RecruiterAddress + ":" + s.CreatorAddress)
	if err != nil {
		return []*schema.Deal{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(deal)), &deals)
	if err != nil {
		return []*schema.Deal{}, err
	}
	return deals, nil
}

func (s *DealController) SetDeal(deals []*schema.Deal) error {
	data, err := json.Marshal(deals)
	if err != nil {
		return err
	}
	fmt.Println("data: ", data)
	record_id := "deals:" + s.RecruiterAddress + ":" + s.CreatorAddress
	fmt.Println("record_id: ", record_id)
	err = repository.JSONWrite("deals:"+s.RecruiterAddress+":"+s.CreatorAddress, data, 0)
	if err != nil {
		return err
	}
	return nil
}
