package web3

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/wealdtech/go-ens/v3"

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

	client, err := ethclient.Dial(conf.Network.Binance.RPCURL)
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
	client.Close()
	return int(state.Int64())
}

func CheckOutstandingPayment(user_address string, token_address string, amount *big.Int, tx_hash string) error {
	conf, err := config.ParseConfig()
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(conf.Network.Binance.RPCURL)
	if err != nil {
		return err
	}

	payment_address := common.HexToAddress(conf.ContractAddresses.JobPayments)
	txHash := common.HexToHash(tx_hash)
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return err
	}
	if isPending {
		return fmt.Errorf("tx is pending")
	}

	if *tx.To() != payment_address {
		return fmt.Errorf("tx to address mismatch")
	}

	//todo: implement multichain payments (currently only binance)
	if tx.ChainId().Int64() != conf.Network.Binance.ID {
		return fmt.Errorf("tx chain id mismatch")
	}

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}

	event_sig := []byte("PaymentAdded(address,uint256)")
	event_sig_hash := crypto.Keccak256Hash(event_sig)

	// block_hash := receipt.BlockHash
	query := ethereum.FilterQuery{
		FromBlock: receipt.BlockNumber,
		ToBlock:   receipt.BlockNumber,
		// BlockHash: &block_hash,
		Addresses: []common.Address{
			payment_address,
		},
		Topics: [][]common.Hash{
			{
				event_sig_hash,
			},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}
	if (logs == nil) || (len(logs) == 0) {
		return fmt.Errorf("no logs found")
	}

	contract_abi, err := abi.JSON(strings.NewReader(string(job_listing.JobListingABI)))
	if err != nil {
		log.Fatal(err)
	}

	//todo: multiple events scenario
	for _, v_log := range logs {
		payment_event, err := contract_abi.Unpack("PaymentAdded", v_log.Data)
		if err != nil {
			return err
		}

		payment_amount := payment_event[0].(*big.Int)
		if payment_amount.Cmp(amount) != 0 {
			return fmt.Errorf("amount paid mismatch")
		}

		token_addr := common.BytesToAddress(v_log.Topics[1].Bytes()).Hex()
		if token_addr != token_address {
			return fmt.Errorf("token address mismatch")
		}

	}
	return nil
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
		highlight_fee = big.NewInt(0)
	}

	var service_fee = new(big.Int)
	service_fee.SetString(conf.Settings.Jobs.ServiceFee, 10)

	total_fee := new(big.Int)
	total_fee.Add(highlight_fee, service_fee)

	return total_fee, nil
}

func CheckNFTOwner(user_address string, token_address string, token_id int) bool {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	client, err := ethclient.Dial(conf.Network.Binance.RPCURL)
	if err != nil {
		log.Fatal(err)
	}

	nft_address_hex := common.HexToAddress(token_address)

	instance, err := genesis.NewGenesis(nft_address_hex, client)
	if err != nil {
		log.Fatal(err)
	}

	user_address_hex := common.HexToAddress(user_address)
	owner, err := instance.OwnerOf(nil, big.NewInt(int64(token_id)))
	if err != nil {
		return false
	}

	if owner.Hex() != user_address_hex.Hex() {
		return false
	}
	return true
}

func CheckENSOwner(user_address string, ens_name string) bool {
	conf, err := config.ParseConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	client, err := ethclient.Dial(conf.Network.Binance.RPCURL)
	if err != nil {
		log.Fatal(err)
	}

	address, err := ens.Resolve(client, ens_name)
	if err != nil {
		return false
	}
	if address.String() != user_address {
		return false
	}
	return true
}
