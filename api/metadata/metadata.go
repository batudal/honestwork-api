package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/takez0o/honestwork-api/utils/web3"
)

type Metadata struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	ExternalUrl string        `json:"external_url"`
	Attributes  []interface{} `json:"attributes"`
}

type TierAttribute struct {
	TraitType string `json:"trait_type"`
	Value     int    `json:"value"`
	MaxValue  int    `json:"max_value"`
}

type GrossRevenueAttribute struct {
	TraitType string `json:"trait_type"`
	Value     int    `json:"value"`
}

type RevenueTierAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Revenue struct {
	NetworkId   int    `json:"network_id"`
	TokenId     int    `json:"token_id"`
	Amount      int    `json:"amount"`
	Tier        int    `json:"tier"`
	RevenueTier string `json:"revenue_tier"`
}

// func fetchAllRevenues() []Revenue {
// 	// get total supply
// 	// scan through revenues with routines

// 	return []Revenue{}
// }

func fetchRevenue(network_id int, token_id int) Revenue {

	amount := web3.FetchNFTRevenue(network_id, token_id)
	fmt.Println("Amount:", amount)
	tier := web3.FetchTokenTier(token_id)
	revenue_tier := getRevenueTier(amount)
	fmt.Println("RevTier:", revenue_tier)

	revenue := Revenue{}
	revenue.NetworkId = network_id
	revenue.TokenId = token_id
	revenue.Amount = amount
	revenue.Tier = tier
	revenue.RevenueTier = revenue_tier
	return revenue
}

func getRevenueTier(amount int) string {
	revenueTiers := []string{
		"< $1000",
		"$1000 - $10,000",
		"$10,000 - $100,000",
		"HonestChad",
	}
	if amount < 1000 {
		return revenueTiers[0]
	} else if amount < 10000 {
		return revenueTiers[1]
	} else if amount < 100000 {
		return revenueTiers[2]
	} else {
		return revenueTiers[3]
	}
}

func main() {
	revenue := fetchRevenue(137, 1)
	fmt.Println("Revenue amount: ", revenue.Amount)
	fmt.Println("Revenue tier: ", revenue.Tier)
	fmt.Println("Revenue tier: ", revenue.RevenueTier)

	writeJSON(revenue)
}

func writeJSON(revenue Revenue) {
	data := Metadata{
		Name:        "HonestWork #" + strconv.Itoa(revenue.TokenId),
		Description: "HonestWork Genesis NFTs are the gateway to HonestWork ecosystem.",
		Image:       "https://honestwork-userfiles.fra1.cdn.digitaloceanspaces.com/genesis-nft/" + strconv.Itoa(revenue.TokenId) + ".png",
		ExternalUrl: "https://honestwork.app",
		Attributes: []interface{}{
			TierAttribute{
				TraitType: "Tier",
				Value:     revenue.Tier,
				MaxValue:  3,
			},
			GrossRevenueAttribute{
				TraitType: "Gross Revenue",
				Value:     revenue.Amount,
			},
			RevenueTierAttribute{
				TraitType: "Revenue Tier",
				Value:     revenue.RevenueTier,
			}}}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(fmt.Sprintf("./static/metadata/%v", revenue.TokenId), file, 0644)
}
