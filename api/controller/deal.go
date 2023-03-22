package controller

import (
	"encoding/json"
	"fmt"

	"github.com/takez0o/honestwork-api/api/repository"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
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
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "GetDeals - JSONRead")
		return []*schema.Deal{}, err
	}
	err = json.Unmarshal([]byte(fmt.Sprint(deal)), &deals)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "GetDeals - Unmarshal")
		return []*schema.Deal{}, err
	}
	return deals, nil
}

func (s *DealController) SetDeal(deals []*schema.Deal) error {
	data, err := json.Marshal(deals)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "SetDeal - Marshal")
		return err
	}
	err = repository.JSONWrite("deals:"+s.RecruiterAddress+":"+s.CreatorAddress, data, 0)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "SetDeal - JSONWrite")
		return err
	}
	return nil
}
