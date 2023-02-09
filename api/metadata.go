package main

import (

)

type Metadata struct {
  Name        string      `json:"name"`
  Description string      `json:"description"`
  Image       string      `json:"image"`
  ExternalUrl string      `json:"external_url"`
  Attributes  []Attribute `json:"attributes"`
}

type Attribute struct {
  TraitType string `json:"trait_type"`
  Value     string `json:"value"`
} 

func fetchRevenue(network_id int, token_id int) int {  
  return 0
}
