package web3

import (
	"fmt"
	"log"
	"math/big"

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

	client, err := ethclient.Dial(conf.Network.Devm.RPCURL)
	if err != nil {
		log.Fatal(err)
	}

	nft_address_hex := common.HexToAddress(conf.ContractAddresses.MembershipNFT)
	fmt.Println("nft_address_hex:", nft_address_hex)

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

func CheckOutstandingPayment(user_address string, token_address string, amount *big.Int) (*job_listing.JobListingPayment, error) {
	conf, err := config.ParseConfig()
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	client, err := ethclient.Dial(conf.Network.Polygon.RPCURL)
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	payment_address_hex := common.HexToAddress(conf.ContractAddresses.JobPayments)
	instance, err := job_listing.NewJobListing(payment_address_hex, client)
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	user_address_hex := common.HexToAddress(user_address)
	outstanding_payment, err := instance.GetLatestPayment(nil, user_address_hex)
	if err != nil {
		return new(job_listing.JobListingPayment), err
	}

	token_address_onchain := outstanding_payment.Token.String()
	amount_is_eq := amount.Cmp(outstanding_payment.Amount)

	if token_address != token_address_onchain || amount_is_eq != 0 {
		return new(job_listing.JobListingPayment), fmt.Errorf("payment mismatch")
	}
	return &outstanding_payment, err
}

func CalculatePayment(opts *schema.Job) (*big.Int, error) {
	conf, err := config.ParseConfig()
	if err != nil {
		return big.NewInt(0), err
	}

	var highlight_fee = new(big.Int)
	duration := opts.StickyDuration
	if duration == 7 {
		highlight_fee.SetString(conf.Settings.Jobs.HighlightPrices.StickyPrices.Week, 10)
	} else if duration == 14 {
		highlight_fee.SetString(conf.Settings.Jobs.HighlightPrices.StickyPrices.Biweek, 10)
	} else if duration == 30 {
		highlight_fee.SetString(conf.Settings.Jobs.HighlightPrices.StickyPrices.Month, 10)
	} else {
		return big.NewInt(0), fmt.Errorf("invalid duration")
	}

	var service_fee = new(big.Int)
	service_fee.SetString(conf.Settings.Jobs.ServiceFee, 10)

	total_fee := new(big.Int)
	total_fee.Add(highlight_fee, service_fee)

	return total_fee, nil
}
