package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func HandleGetDeals() fiber.Handler {
	return func(c *fiber.Ctx) error {
		deal_controller := controller.NewDealController(c.Params("recruiter"), c.Params("creator"))
		deals, err := deal_controller.GetDeals()
		if err != nil {
			return c.JSON([]*schema.Deal{})
		}
		return c.JSON(deals)
	}
}

func HandleAddDeal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		deal_controller := controller.NewDealController(c.Params("recruiter"), c.Params("creator"))
		deals, _ := deal_controller.GetDeals()

		var deal *schema.Deal
		err := json.Unmarshal(c.Body(), &deal)
		if err != nil {
			return err
		}
		deal.Status = "offered"
		// todo: check if given job id exists and not consumed

		deals = append(deals, deal)
		err = deal_controller.SetDeal(deals)
		if err != nil {
			return err
		}

		return c.JSON("success")
	}
}

func HandleSignDeal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		deal_controller := controller.NewDealController(c.Params("recruiter"), c.Params("creator"))
		deals, err := deal_controller.GetDeals()
		if err != nil {
			return err
		}

		type DealSignature struct {
			Slot      int    `json:"slot"`
			Signature string `json:"signature"`
		}

		var dealSignature DealSignature
		err = json.Unmarshal(c.Body(), &dealSignature)
		if err != nil {
			return err
		}

		if dealSignature.Slot > len(deals) {
			return fiber.NewError(fiber.StatusBadRequest, "Wrong slot.")
		}

		deals[dealSignature.Slot].Signature = dealSignature.Signature
		deals[dealSignature.Slot].Status = "accepted"

		err = deal_controller.SetDeal(deals)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}

func HandleExecuteDeal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		deal_controller := controller.NewDealController(c.Params("recruiter"), c.Params("creator"))
		deals, err := deal_controller.GetDeals()
		if err != nil {
			return err
		}

		type DealExecution struct {
			Slot int `json:"slot"`
		}

		var dealExecution DealExecution
		err = json.Unmarshal(c.Body(), &dealExecution)
		if err != nil {
			return err
		}

		if dealExecution.Slot > len(deals) {
			return fiber.NewError(fiber.StatusBadRequest, "Wrong slot.")
		}

		deals[dealExecution.Slot].Status = "executed"

		err = deal_controller.SetDeal(deals)
		if err != nil {
			return err
		}
		return c.JSON("success")
	}
}
