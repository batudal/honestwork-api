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
			ServiceFee   string `yaml:"service_fee"`
			StickyPrices struct {
				Week   string `yaml:"7_days"`
				Biweek string `yaml:"14_days"`
				Month  string `yaml:"30_days"`
			} `yaml:"sticky_prices"`
		} `yaml:"jobs"`
	} `yaml:"settings"`
	ContractAddresses struct {
		MembershipNFT string `yaml:"membership_nft"`
		JobPayments   string `yaml:"job_payments"`
		Registry      string `yaml:"registry"`
		Escrow        string `yaml:"escrow"`
	} `yaml:"contract_addresses"`
	Network struct {
		Eth struct {
			ID int `yaml:"id"`
		} `yaml:"eth"`
		Arbitrum struct {
			ID int `yaml:"id"`
		} `yaml:"arbitrum"`
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
