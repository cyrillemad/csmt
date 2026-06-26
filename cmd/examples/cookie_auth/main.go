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
			steam.Cookie{
				SteamLoginSecure: "",
				SessionID:        "",
			}))
	/*
		some methods require cookie authorization
		like as steamcommunity history, right now
		you need manually paste it from your browser
	*/
	//todo: auto cookie generating (with steam-auth maybe)

	hash, err := client.Community.SearchHash("Dreams and nightmares")
	if err != nil {
		fmt.Println(err)
	}
	history, err := client.Community.PriceHistory(hash)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range history {
		fmt.Println(entry.Time)
		fmt.Println(entry.Price)
	}
}
