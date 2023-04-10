package worker

import (
	"context"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/takez0o/honestwork-api/api/controller"
	"github.com/takez0o/honestwork-api/utils/abi/hwescrow"
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
	client, err := ethclient.Dial(os.Getenv("ARBITRUM_WSS"))
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

	escrow_abi, err := abi.JSON(strings.NewReader(string(hwescrow.HwescrowABI)))
	if err != nil {
		log.Fatal(err)
	}
	type OfferEvent struct {
		Recruiter    common.Address
		Creator      common.Address
		TotalPayment *big.Int
		PaymentToken common.Address
		JobId        *big.Int
	}

	logOfferCreatedSig := []byte("OfferCreated(address,address,uint256,address,uint256)")
	logOfferCreatedHash := crypto.Keccak256Hash(logOfferCreatedSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-logs:
			if log.Topics[0] == logOfferCreatedHash {
				var offerEvent OfferEvent
				offerEvent.Recruiter = common.HexToAddress(log.Topics[1].Hex())
				_ = escrow_abi.UnpackIntoInterface(&offerEvent, "OfferCreated", log.Data)
				updateJob(conf, offerEvent.Recruiter.String(), int(offerEvent.JobId.Int64()))
			}
		}
	}
}

func updateJob(conf *config.Config, address string, slot int) {
	job_controller := controller.NewJobController(address, slot)
	job, err := job_controller.GetJob()
	if err != nil {
		log.Fatal(err)
	}
	if job.UserAddress == address && job.DealId == -1 {
		job.DealId = slot
		job.DealNetworkId = conf.Network.Arbitrum.ID
		job_controller.SetJob(&job)
	}
}
