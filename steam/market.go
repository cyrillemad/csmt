package steam

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cyrillemad/csmt/internal/encode"
	"github.com/cyrillemad/csmt/types"
)

func (steam *Client) getPriceOverview(
	ctx context.Context,
	hash types.MarketHash,
	v any) error {

	query := url.Values{}

	query.Set("country", steam.config.Country)
	query.Set("appid", strconv.Itoa(steam.config.AppID))
	query.Set("currency", strconv.Itoa(int(steam.config.Currency)))
	query.Set("market_hash_name", string(hash))

	path := "/market/priceoverview/?" + query.Encode()

	err := steam.Client.Get(ctx, path, v)

	if err != nil {
		return err
	}

	return nil
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
	result.Volume, err = strconv.Atoi(response.Volume)
	if err != nil {
		return result, err
	}

	return result, nil
}
