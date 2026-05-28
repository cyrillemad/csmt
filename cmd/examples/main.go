package main

import (
	"fmt"

	"github.com/cyrillemad/csmt/steam"
	"github.com/cyrillemad/csmt/types"
)

const (
	TestHash types.MarketHash = "AK-47 | Redline (Field-Tested)"
)

func main() {
	client := steam.NewClient(steam.WithCurrency(types.USD))
	price, err := client.PriceOverview(TestHash)
	if err != nil {
		panic(err)
	}
	fmt.Println(price.LowestPrice, "USD") //todo: create maps for steam enums

	hashes, err := client.SearchHash("Case")
	if err != nil {
		panic(err)
	}
	for _, hash := range hashes {
		fmt.Println(hash)
	}
}
