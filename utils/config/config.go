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
	}

	DB struct {
		Port string `yaml:"port"`
		ID   int    `yaml:"id"`
	}

	Settings struct {
		Skills      Tiers
		Jobs        Tiers
		Duration    Tiers
		Referral    Tiers
		Price       Tiers
		Application Tiers
	}
}

func ParseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
