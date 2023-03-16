package worker

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	fetchDeals()
	time.Sleep(time.Duration(30) * time.Minute)
}

func fetchDeals() {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	client, err := ethclient.Dial(conf.Network.Arbitrum.RPCURL)
	if err != nil {
		log.Fatal(err)
	}

	escrow_address_hex := common.HexToAddress(conf.ContractAddresses.Escrow)

	instance, err := hwescrow.NewHwescrow(escrow_address_hex, client)
	if err != nil {
		log.Fatal(err)
	}

	deals, err := instance.GetAllDeals(nil)
	for _, deal := range deals {
		writeJob(&deal)

	}

	defer client.Close()

}

func writeJob(deal *hwescrow.HWEscrowDeal) {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	jobs_controller := controller.NewJobIndexer("jobsIndex")
	jobs, err := jobs_controller.GetAllJobs()
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, job := range jobs {
		if job.DealId == -1 {
			if job.UserAddress == addressToString(deal.Recruiter) {
				if job.Slot == toInt(deal.JobId) {
					job_writer := controller.NewJobController(job.UserAddress, job.Slot)
					job.DealId = toInt(deal.JobId)
					job.DealNetworkId = int(conf.Network.Arbitrum.ID)
					job_writer.SetJob(&job)
				}
			}
		}
	}

}

func toBigInt(num int) *big.Int {
	return big.NewInt(int64(num))
}

func toInt(num *big.Int) int {
	return int(num.Int64())
}

func addressToString(address common.Address) string {
	return address.Hex()
}
