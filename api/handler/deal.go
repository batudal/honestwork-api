package handler

import (
	"encoding/json"

	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/loggersentry"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetDeals(recruiter string, creator string) []*schema.Deal {
	deal_controller := controller.NewDealController(recruiter, creator)
	deals, err := deal_controller.GetDeals()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleGetDeals")
		return []*schema.Deal{}
	}
	return deals
}

func HandleAddDeal(recruiter string, creator string, signature string, body []byte) string {
	deal_controller := controller.NewDealController(recruiter, creator)
	deals, _ := deal_controller.GetDeals()

	var deal *schema.Deal
	err := json.Unmarshal(body, &deal)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleAddDeal")
		return err.Error()
	}
	deal.Status = "offered"
	// todo: check if given job id exists and not consumed

	deals = append(deals, deal)
	err = deal_controller.SetDeal(deals)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleAddDeal")
		return err.Error()
	}

	return "success"
}

func HandleSignDeal(recruiter string, creator string, signature string, body []byte) string {
	deal_controller := controller.NewDealController(recruiter, creator)
	deals, err := deal_controller.GetDeals()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleSignDeal")
		return err.Error()
	}

	type DealSignature struct {
		Slot      int    `json:"slot"`
		Signature string `json:"signature"`
	}

	var dealSignature DealSignature
	err = json.Unmarshal(body, &dealSignature)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleSignDeal")
		return err.Error()
	}

	if dealSignature.Slot > len(deals) {
		return "Wrong slot."
	}

	deals[dealSignature.Slot].Signature = dealSignature.Signature
	deals[dealSignature.Slot].Status = "accepted"

	err = deal_controller.SetDeal(deals)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleSignDeal")
		return err.Error()
	}
	return "success"
}

func HandleExecuteDeal(recruiter string, creator string, signature string, body []byte) string {
	deal_controller := controller.NewDealController(recruiter, creator)
	deals, err := deal_controller.GetDeals()
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleExecuteDeal")
		return err.Error()
	}

	type DealExecution struct {
		Slot int `json:"slot"`
	}

	var dealExecution DealExecution
	err = json.Unmarshal(body, &dealExecution)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleExecuteDeal")
		return err.Error()
	}

	if dealExecution.Slot > len(deals) {
		return "Wrong slot."
	}

	deals[dealExecution.Slot].Status = "executed"

	err = deal_controller.SetDeal(deals)
	if err != nil {
		loggersentry.InitSentry()
		loggersentry.CaptureErrorMessage(err.Error() + "handleExecuteDeal")
		return err.Error()
	}
	return "success"
}
