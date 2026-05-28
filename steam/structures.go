package steam

import "github.com/cyrillemad/csmt/types"

type PriceOverviewResponse struct {
	Status      bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	MedianPrice string `json:"median_price"`
	Volume      string `json:"volume"`
}

type PriceOverview struct {
	LowestPrice float64        `json:"lowest_price"`
	MedianPrice float64        `json:"median_price"`
	Volume      int            `json:"volume"`
	Currency    types.Currency `json:"currency"`
}
