package web3

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/getsentry/sentry-go"
	"github.com/wealdtech/go-ens/v3"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/takez0o/honestwork-api/utils/abi/honestworknft"
	"github.com/takez0o/honestwork-api/utils/abi/hwescrow"
	"github.com/takez0o/honestwork-api/utils/abi/hwlisting"
	"github.com/takez0o/honestwork-api/utils/abi/hwregistry"
	"github.com/takez0o/honestwork-api/utils/config"
	"github.com/takez0o/honestwork-api/utils/schema"
)

func FetchUserNFT(address string) int {
	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureException(err)
	}

	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))
	if err != nil {
		sentry.CaptureException(err)
	}
	defer client.Close()

	nft_address_hex := common.HexToAddress(conf.ContractAddresses.MembershipNFT)

	instance, err := honestworknft.NewHonestworknft(nft_address_hex, client)
	if err != nil {
		sentry.CaptureException(err)
	}

	user_address_hex := common.HexToAddress(address)
	index := big.NewInt(0)
	token_id, err := instance.TokenOfOwnerByIndex(nil, user_address_hex, index)
	if err != nil {
		sentry.CaptureException(err)
	}
	return int(token_id.Int64())
}

func FetchUserState(address string) int {
	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureException(err)
	}

	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))
	if err != nil {
		sentry.CaptureException(err)
	}

	nft_address_hex := common.HexToAddress(conf.ContractAddresses.MembershipNFT)

	instance, err := honestworknft.NewHonestworknft(nft_address_hex, client)
	if err != nil {
		sentry.CaptureException(err)
	}

	user_address_hex := common.HexToAddress(address)
	state, err := instance.GetUserTier(nil, user_address_hex)
	if err != nil {
		sentry.CaptureException(err)
	}
	client.Close()
	return int(state.Int64())
}

func FetchTokenTier(token_id int) int {
	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureException(err)
	}

	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))
	if err != nil {
		sentry.CaptureException(err)
	}

	nft_address_hex := common.HexToAddress(conf.ContractAddresses.MembershipNFT)

	instance, err := honestworknft.NewHonestworknft(nft_address_hex, client)
	if err != nil {
		sentry.CaptureException(err)
	}

	token_id_big := big.NewInt(int64(token_id))
	state, err := instance.GetTokenTier(nil, token_id_big)
	if err != nil {
		sentry.CaptureException(err)
	}
	client.Close()
	return int(state.Int64())
}

func FetchNFTRevenue(network_id int, token_id int) int {
	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureException(err)
	}

	var client *ethclient.Client
	if network_id == 42161 {
		client, err = ethclient.Dial(os.Getenv("ARBITRUM_RPC"))
		if err != nil {
			sentry.CaptureException(err)
		}
	}
	defer client.Close()

	registry_address_hex := common.HexToAddress(conf.ContractAddresses.Registry)
	instance, err := hwregistry.NewHwregistry(registry_address_hex, client)
	if err != nil {
		sentry.CaptureException(err)
	}

	revenue, err := instance.GetNFTGrossRevenue(nil, big.NewInt(int64(token_id)))
	if err != nil {
		sentry.CaptureException(err)
	}

	revenue_normalized := new(big.Int)
	revenue_normalized.Div(revenue, big.NewInt(1000000000000000000))
	return int(revenue_normalized.Int64())
}

func FetchTotalSupply() int {
	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureException(err)
	}

	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))
	if err != nil {
		sentry.CaptureException(err)
	}

	nft_address_hex := common.HexToAddress(conf.ContractAddresses.MembershipNFT)
	instance, err := honestworknft.NewHonestworknft(nft_address_hex, client)
	if err != nil {
		sentry.CaptureException(err)
	}

	total_supply, err := instance.TotalSupply(nil)
	if err != nil {
		sentry.CaptureException(err)
	}
	client.Close()
	return int(total_supply.Int64())
}

func CheckOutstandingPayment(user_address string, token_address string, amount *big.Int, tx_hash string) error {
	conf, err := config.ParseConfig()
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(os.Getenv("ARBITRUM_RPC"))
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

	if int(tx.ChainId().Int64()) != conf.Network.Arbitrum.ID {
		return fmt.Errorf("tx chain id mismatch")
	}

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return err
	}

	event_sig_hash := crypto.Keccak256Hash([]byte("PaymentAdded(address,uint256)"))

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

	contract_abi, err := abi.JSON(strings.NewReader(string(hwlisting.HwlistingABI)))
	if err != nil {
		sentry.CaptureException(err)
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

	var extra = new(big.Int)
	duration := opts.StickyDuration
	if duration == 7 {
		extra.SetString(conf.Settings.Jobs.StickyPrices.Week, 10)
	} else if duration == 14 {
		extra.SetString(conf.Settings.Jobs.StickyPrices.Biweek, 10)
	} else if duration == 30 {
		extra.SetString(conf.Settings.Jobs.StickyPrices.Month, 10)
	} else {
		extra = big.NewInt(0)
	}

	var service_fee = new(big.Int)
	service_fee.SetString(conf.Settings.Jobs.ServiceFee, 10)

	total_fee := new(big.Int)
	total_fee.Add(extra, service_fee)

	return total_fee, nil
}

func CheckNFTOwner(user_address string, token_address string, token_id int) bool {
	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))
	if err != nil {
		sentry.CaptureException(err)
	}

	nft_address_hex := common.HexToAddress(token_address)

	instance, err := honestworknft.NewHonestworknft(nft_address_hex, client)
	if err != nil {
		sentry.CaptureException(err)
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
	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))
	if err != nil {
		sentry.CaptureException(err)
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

func FetchAggregatedRating(user_address string) float64 {
	conf, err := config.ParseConfig()
	if err != nil {
		sentry.CaptureException(err)
	}

	var client *ethclient.Client
	client, err = ethclient.Dial(os.Getenv("ARBITRUM_RPC"))
	if err != nil {
		sentry.CaptureException(err)
	}
	defer client.Close()

	escrow_address_hex := common.HexToAddress(conf.ContractAddresses.Escrow)
	instance, err := hwescrow.NewHwescrow(escrow_address_hex, client)
	if err != nil {
		sentry.CaptureException(err)
	}

	user_address_hex := common.HexToAddress(user_address)
	rating, err := instance.GetAggregatedRating(nil, user_address_hex)
	if err != nil {
		sentry.CaptureException(err)
	}

	precision, err := instance.GetPrecision(nil)
	if err != nil {
		sentry.CaptureException(err)
	}
	rating_normalized := float64(rating.Int64()) / float64(precision.Int64())
	return rating_normalized
}
