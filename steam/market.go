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

func (steam *Client) Price(
	hash types.MarketHash) (float64, error) {

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
	if err != nil {
		return -1, err
	}
	if response.Status != true {
		return -1, fmt.Errorf("Steam returned non-success response")
	}
	return encode.ParsePrice(response.LowestPrice)
}
