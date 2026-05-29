package steamcommunity

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cyrillemad/csmt/internal/encode"
	"github.com/cyrillemad/csmt/types"
)

type PriceOverview struct {
	LowestPrice float64        `json:"lowest_price"`
	MedianPrice float64        `json:"median_price"`
	Volume      int            `json:"volume"`
	Currency    types.Currency `json:"currency"`
}

func (steam *Client) PriceOverview(
	hash types.MarketHash) (PriceOverview, error) {

	ctx := context.Background()

	if steam.config.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(
			ctx,
			steam.config.Timeout,
		)

		defer cancel()
	}

	response := PriceOverviewResponse{}
	err := steam.getPriceOverview(ctx, hash, &response)
	result := PriceOverview{}

	if err != nil {
		return result, err
	}
	if response.Status != true {
		return result, fmt.Errorf("Steam returned non-success response")
	}

	result.Currency = steam.config.Currency
	result.LowestPrice, err = encode.ParsePrice(response.LowestPrice)
	if err != nil {
		return result, err
	}
	result.MedianPrice, err = encode.ParsePrice(response.MedianPrice)
	if err != nil {
		return result, err
	}
	result.Volume, err = strconv.Atoi(
		strings.ReplaceAll(response.Volume, ",", ""))
	if err != nil {
		return result, err
	}
	return result, nil
}

func (steam *Client) SearchHashQuery(
	name string) (hashes []types.MarketHash, err error) {
	ctx := context.Background()

	if steam.config.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(
			ctx,
			steam.config.Timeout,
		)

		defer cancel()
	}

	response := RenderSearchResponse{}
	request := RenderSearchOptions{
		Query:         name,
		SortColumn:    SortColumnName,
		SortDirection: SortAsc,
		NoRender:      true,
	}
	err = steam.getRenderSearch(ctx, request, &response)
	if err != nil {
		return hashes, err
	}
	for _, result := range response.Results {
		hashes = append(hashes, types.MarketHash(result.HashName))
	}
	return hashes, nil
}

func (steam *Client) SearchHash(
	name string) (hash types.MarketHash, err error) {
	hashes, err := steam.SearchHashQuery(name)
	if len(hashes) > 0 {
		return hashes[0], err
	}
	return hash, fmt.Errorf("Steam returned no results")
}
