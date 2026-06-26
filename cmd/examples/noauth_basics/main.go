package main

import (
	"fmt"

	"github.com/cyrillemad/csmt"
	steam "github.com/cyrillemad/csmt/steamcommunity"
	"github.com/cyrillemad/csmt/types"
)

func main() {
	client := csmt.NewNoAuthClient(
		steam.WithCurrency(types.USD))

	hash, err := client.Community.SearchHash("Dreams and nightmares")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)

	price, err := client.Community.PriceOverview(hash)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(price.MedianPrice, "USD")

	url, err := client.Apis.HashImageURL(hash)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)

	items, err := client.Community.Inventory("76561199416551019")
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		fmt.Println(item.Name, item.Amount)
	}
}
