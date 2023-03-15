package worker

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/config"
)

type EventSubscriber struct {
}

type Offer struct {
	_recruiter    common.Hash
	_creator      common.Hash
	_totalPayment common.Hash
	_paymentToken common.Hash
	_jobId        common.Hash
}

func NewEventSubscriber() *EventSubscriber {
	return &EventSubscriber{}
}

func (r *EventSubscriber) Subscribe() {
	connect()
}

func connect() {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	client, err := ethclient.Dial(os.Getenv("ARBITRUM_WEBSOCKET"))
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(conf.ContractAddresses.Escrow)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}
	logOfferCreatedSig := []byte("OfferCreated(address,address,uint256,address,uint256)")
	logOfferCreatedHash := crypto.Keccak256Hash(logOfferCreatedSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-logs:
			if log.Topics[0] == logOfferCreatedHash {
				writeJobsForSubscriber(hexToInt(log.Topics[len(log.Topics)-1]), string(hashToAddress(log.Topics[1]).Hex()))
			}
		}

	}
}

func hexToInt(hex common.Hash) int {
	i, err := strconv.ParseInt(hex.Hex(), 0, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(i)
}

func writeJobsForSubscriber(_jobId int, _address string) {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	jobs_controller := controller.NewJobIndexer("jobsIndex")
	jobs, err := jobs_controller.GetAllJobs()
	if err != nil {
		log.Fatal(err)
	}
	for _, job := range jobs {
		if job.UserAddress == _address {
			if job.DealId == -1 {
				job_writer := controller.NewJobController(job.UserAddress, job.Slot)
				job.DealId = _jobId
				job.DealNetworkId = int(conf.Network.Arbitrum.ID)
				job_writer.SetJob(&job)
			}
		}
	}
}

func hashToAddress(hash common.Hash) common.Address {
	return common.BytesToAddress(hash.Bytes())
}
