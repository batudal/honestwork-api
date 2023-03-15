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
	deal_count := getDealCount()

	chain_map := fetchDeals(deal_count)
	getJobs(chain_map)
	time.Sleep(time.Duration(30) * time.Minute)
}

func fetchDeals(dealAmount int) map[string][]int64 {
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

	//loop through all deals
	addr_map := make(map[string][]int64)
	for i := 0; i <= dealAmount; i++ {
		deal, err := instance.GetDeal(nil, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}
		if deal.Recruiter.String() != "0x0000000000000000000000000000000000000000" {
			addr_map[deal.Recruiter.String()] = append(addr_map[deal.Recruiter.String()], deal.JobId.Int64())
			fmt.Println("addr map: ", addr_map)
		}

	}
	defer client.Close()
	return addr_map

}

func getDealCount() int {
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

	dealCount, err := instance.DealIds(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deal count:", dealCount)

	defer client.Close()
	return int(dealCount.Int64())
}

func getJobs(chain_map map[string][]int64) {
	//get all skills and loop
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}
	jobs_controller := controller.NewJobIndexer("jobsIndex")
	jobs, err := jobs_controller.GetAllJobs()
	if err != nil {
		log.Fatal(err)
	}
	for recruiterAddresses, jobIds := range chain_map {
		fmt.Println("key:", recruiterAddresses, "element:", jobIds)
		for _, job := range jobs {
			if job.UserAddress == recruiterAddresses {
				if job.DealId == -1 {

					job_writer := controller.NewJobController(job.UserAddress, job.Slot)
					job.DealId = int(jobIds[len(jobIds)-1])
					job.DealNetworkId = int(conf.Network.Arbitrum.ID)
					job_writer.SetJob(&job)
				}
			}
		}

	}

	// jobs[0].DealId = "10"
	// job_writer := controller.NewJobController(jobs[0].UserAddress, jobs[0].Slot)
	// job_writer.SetJob(&jobs[0])
	//set deal id for each job in database

}

// checks instead

// get network_id, recruiter_addr + job_id (job:recruiter_addr:job_id)

//JOB DEFINITION: map[recruiteraddress] = job id
// check if job has already been updated
// update job on database with the deal_network_id and deal_id

//update contract -- add timestamp to deal struct -- add random event to check --check getAllDeals
//add subscription to deal watcher
