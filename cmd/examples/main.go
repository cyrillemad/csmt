package main

import (
	"fmt"

	"github.com/cyrillemad/csmt/steamapis"
	steam "github.com/cyrillemad/csmt/steamcommunity"
	"github.com/cyrillemad/csmt/types"
)

const (
	TestHash types.MarketHash = "AK-47 | Redline (Field-Tested)"
)

func main() {
	client := steam.NewClient(steam.WithCurrency(types.USD))
	sapi := steamapis.NewClient() // make united client for no-auth apis

	hash, err := client.SearchHash("Dreams and nightmares")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash)

	price, err := client.PriceOverview(hash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(price.MedianPrice, "USD") //todo: create maps for steamcommunity enums

	url, err := sapi.HashImageURL(hash)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)
}
