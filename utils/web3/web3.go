package web3

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	genesis "github.com/takez0o/honestwork-api/utils/abi"
)

func FetchUserState(address string) int {
	nft_address := "0x32058e2CCdAA0b4615994362d44cC64dFFd3340A"

	client, err := ethclient.Dial("https://decoded.wtf")
	if err != nil {
		log.Fatal(err)
	}

	nft_address_hex := common.HexToAddress(nft_address)
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
