package main

import (
	"fmt"

	"github.com/cyrillemad/csmt"
	"github.com/cyrillemad/csmt/types"
)

const (
	TestHash types.MarketHash = "AK-47 | Redline (Field-Tested)"
)

func main() {
	client := csmt.NewNoAuthClient()

	hash, err := client.Community.SearchHash("Dreams and nightmares")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash)

	price, err := client.Community.PriceOverview(hash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(price.MedianPrice, "USD") //todo: create maps for steamcommunity enums

	url, err := client.Apis.HashImageURL(hash)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)
}
