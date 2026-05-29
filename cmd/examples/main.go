package main

import (
	"fmt"

	steam "github.com/cyrillemad/csmt/steamcommunity"
	"github.com/cyrillemad/csmt/types"
)

const (
	TestHash types.MarketHash = "AK-47 | Redline (Field-Tested)"
)

func main() {
	client := steam.NewClient(steam.WithCurrency(types.USD))

	hash, err := client.SearchHash("Dreams and nightmares")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash)

	price, err := client.PriceOverview(hash)
	if err != nil {
		panic(err)
	}
	fmt.Println(price.MedianPrice, "USD") //todo: create maps for steamcommunity enums
}
