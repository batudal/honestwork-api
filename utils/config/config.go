package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Tiers struct {
	Tier_1 int `yaml:"tier_1"`
	Tier_2 int `yaml:"tier_2"`
	Tier_3 int `yaml:"tier_3"`
}

type Config struct {
	API struct {
		Port string `yaml:"port"`
	} `yaml:"api"`
	DB struct {
		Port string `yaml:"port"`
		ID   int    `yaml:"id"`
	} `yaml:"db"`
	Settings struct {
		Skills Tiers
		Jobs   struct {
			ServiceFee         string `yaml:"service_fee"`
			SuccessFeeBrackets struct {
				Bracket_0 int `yaml:"bracket_0"`
				Bracket_1 int `yaml:"bracket_1"`
				Bracket_2 int `yaml:"bracket_2"`
				Bracket_3 int `yaml:"bracket_3"`
				Bracket_4 int `yaml:"bracket_4"`
			} `yaml:"success_fee_brackets"`
			SuccessFeePercentages struct {
				Percentage_0 int `yaml:"percentage_0"`
				Percentage_1 int `yaml:"percentage_1"`
				Percentage_2 int `yaml:"percentage_2"`
				Percentage_3 int `yaml:"percentage_3"`
				Percentage_4 int `yaml:"percentage_4"`
			} `yaml:"success_fee_percentages"`
			HighlightPrices struct {
				StickyPrices struct {
					Week   string `yaml:"7_days"`
					Biweek string `yaml:"14_days"`
					Month  string `yaml:"30_days"`
				} `yaml:"sticky_prices"`
				HighlightFrame struct {
					Default string `yaml:"default"`
				} `yaml:"highlight_frame"`
			} `yaml:"highlight_prices"`
		} `yaml:"jobs"`
		CharLimits struct {
			Profile struct {
				Username int `yaml:"username"`
				Title    int `yaml:"title"`
				Bio      int `yaml:"bio"`
			} `yaml:"profile"`
		} `yaml:"char_limits"`
	} `yaml:"settings"`
	ContractAddresses struct {
		MembershipNFT string `yaml:"membership_nft"`
		JobPayments   string `yaml:"job_payments"`
	} `yaml:"contract_addresses"`
	Network struct {
		Devm struct {
			ID     int64  `yaml:"id"`
			RPCURL string `yaml:"rpc_url"`
		} `yaml:"devm"`
		Polygon struct {
			ID     int64  `yaml:"id"`
			RPCURL string `yaml:"rpc_url"`
		} `yaml:"polygon"`
		Binance struct {
			ID     int64  `yaml:"id"`
			RPCURL string `yaml:"rpc_url"`
		} `yaml:"binance"`
	} `yaml:"network"`
}

func ParseConfig() (*Config, error) {
	data, err := os.ReadFile("../config.yaml")
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
