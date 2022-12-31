package web3

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/takez0o/honestwork-api/utils/abi/genesis"
	"github.com/takez0o/honestwork-api/utils/abi/job_listing"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func FetchUserState(address string) int {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	client, err := ethclient.Dial(conf.Network.ChainID.RPCURL)
	if err != nil {
		log.Fatal(err)
	}

	nft_address_hex := common.HexToAddress(conf.ContractAddresses.MembershipNFT)
	instance, err := genesis.NewGenesis(nft_address_hex, client)
	if err != nil {
		log.Fatal(err)
	}

	user_address_hex := common.HexToAddress(address)
	state, err := instance.GetUserState(nil, user_address_hex)
	if err != nil {
		log.Fatal(err)
	}

	return int(state.Int64())
}

func CheckOutstandingPayment(user_address string, payment_id []byte, token_address string, amount int64) (*job_listing.JobListingPayment, error) {
	conf, err := config.ParseConfig()
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	client, err := ethclient.Dial(conf.Network.ChainID.RPCURL)
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	payment_address_hex := common.HexToAddress(conf.ContractAddresses.JobPayments)
	instance, err := job_listing.NewJobListing(payment_address_hex, client)
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	user_address_hex := common.HexToAddress(user_address)
	outstanding_payment, err := instance.GetPaymentById(nil, user_address_hex, payment_id)
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	token_address_onchain := outstanding_payment.Token.String()
	if token_address != token_address_onchain || outstanding_payment.Amount.Int64() != amount {
		return new(job_listing.JobListingPayment), fmt.Errorf("payment mismatch")
	}
	return &outstanding_payment, err
}

func CalculatePayment(*schema.HighlightOpts) (int64, error) {
	//todo: implement calculation
	return 0, nil
}
