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
	conf, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
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
		fmt.Println("Error:", err)
	}

	logOfferCreatedSig := []byte("OfferCreated(address,address,uint256,address,uint256)")
	logOfferCreatedHash := crypto.Keccak256Hash(logOfferCreatedSig)

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Error:", err)
		case log := <-logs:
			if log.Topics[0] == logOfferCreatedHash {
				updateJob(string(hashToAddress(log.Topics[1]).Hex()), hexToInt(log.Topics[len(log.Topics)-1]))
			}
		}
	}
}

func hexToInt(hex common.Hash) int {
	i, _ := strconv.ParseInt(hex.Hex(), 0, 64)
	return int(i)
}

func updateJob(address string, slot int) {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	job_controller := controller.NewJobController(address, slot)
	job, err := job_controller.GetJob()
	if err != nil {
		fmt.Println("Error:", err)
	}
	if job.UserAddress == address && job.DealId == -1 {
		job.DealId = slot
		job.DealNetworkId = int(conf.Network.Arbitrum.ID)
		job_controller.SetJob(&job)

	}
}

func hashToAddress(hash common.Hash) common.Address {
	return common.BytesToAddress(hash.Bytes())
}
