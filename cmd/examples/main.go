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
	client := steam.NewClient()
	price, err := client.Price(TestHash)
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
}
