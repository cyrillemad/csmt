package main

import (
	"fmt"

	"github.com/cyrillemad/csmt"
	steam "github.com/cyrillemad/csmt/steamcommunity"
	"github.com/cyrillemad/csmt/types"
)

func main() {
	client := csmt.NewNoAuthClient(
		steam.WithCurrency(types.USD),
		steam.WithCookie(
			"steamLoginSecure=xxx; "+
				"sessionid=xxx"))

	/*
		some methods require cookie authorization
		like as steamcommunity history, right now
		you need manually paste it from your browser
	*/

	//todo: auto cookie generating (with steam-auth maybe)
	//todo: smart cookie system

	hash, err := client.Community.SearchHash("Dreams and nightmares")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", hash)
	history, err := client.Community.PriceHistory(hash)

	fmt.Println(err)
	fmt.Printf("%+v\n", history)
	for _, entry := range history {
		fmt.Println(entry.Time)
		fmt.Println(entry.Price)
	}
}
