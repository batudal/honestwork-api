package worker

import (
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/getsentry/sentry-go"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/abi/hwescrow"
	"github.com/takez0o/honestwork-api/utils/config"
)

type DealWatcher struct {
}

func NewDealWatcher() *DealWatcher {
	return &DealWatcher{}
}

func (r *DealWatcher) WatchDeals() {
	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureException(err)
	}
	client, err := ethclient.Dial(os.Getenv("ARBITRUM_RPC"))
	if err != nil {
		sentry.CaptureException(err)
	}
	defer client.Close()
	for {
		updateDeals(conf, client)
		time.Sleep(time.Duration(4) * time.Hour)
	}
}

func updateDeals(conf *config.Config, client *ethclient.Client) {
	escrow_address_hex := common.HexToAddress(conf.ContractAddresses.Escrow)
	instance, err := hwescrow.NewHwescrow(escrow_address_hex, client)
	if err != nil {
		return
	}
	deals, err := instance.GetDeals(nil)
	if err != nil {
		return
	}
	for i, deal := range deals {
		job_controller := controller.NewJobController(deal.Recruiter.String(), int(deal.JobId.Int64()))
		job, _ := job_controller.GetJob()
		if job.DealId == -1 {
			job_writer := controller.NewJobController(job.UserAddress, job.Slot)
			job.DealId = i
			job.DealNetworkId = int(conf.Network.Arbitrum.ID)
			job_writer.SetJob(&job)
		}
	}
}
