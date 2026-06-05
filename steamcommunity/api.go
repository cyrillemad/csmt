package steamcommunity

import (
	"context"
	"net/url"
	"strconv"

	"fmt"
	"github.com/cyrillemad/csmt/types"
)

func (steam *Client) getPriceOverview(
	ctx context.Context,
	hash types.MarketHash,
	v *PriceOverviewResponse) error {

	query := url.Values{}

	query.Set("country", steam.config.Country)
	query.Set("appid", strconv.Itoa(steam.config.AppID))
	query.Set("currency", strconv.Itoa(int(steam.config.Currency)))
	query.Set("market_hash_name", string(hash))

	path := steam.config.APIPath + "/market/priceoverview/?" + query.Encode()
	err := steam.Client.Get(
		ctx,
		path,
		types.Authorize{},
		v)

	if err != nil {
		return err
	}

	return nil
}

func (steam *Client) getRenderSearch(
	ctx context.Context,
	options RenderSearchOptions,
	v *RenderSearchResponse,
) error {

	query := url.Values{}

	query.Set(
		"appid",
		strconv.Itoa(steam.config.AppID),
	)

	if options.Query != "" {
		query.Set("query", options.Query)
	}

	if options.Start > 0 {
		query.Set(
			"start",
			strconv.Itoa(options.Start),
		)
	}

	if options.Count > 0 {
		query.Set(
			"count",
			strconv.Itoa(options.Count),
		)
	}

	if options.SortColumn != "" {
		query.Set(
			"sort_column",
			string(options.SortColumn),
		)
	}

	if options.SortDirection != "" {
		query.Set(
			"sort_dir",
			string(options.SortDirection),
		)
	}

	if options.SearchDescriptions {
		query.Set(
			"search_descriptions",
			"1",
		)
	}

	if options.PriceMin > 0 {
		query.Set(
			"price_min",
			strconv.Itoa(options.PriceMin),
		)
	}

	if options.PriceMax > 0 {
		query.Set(
			"price_max",
			strconv.Itoa(options.PriceMax),
		)
	}

	if options.NoRender {
		query.Set("norender", "1")
	}

	path := steam.config.APIPath + "/market/search/render/?" + query.Encode()

	err := steam.Client.Get(
		ctx,
		path,
		types.Authorize{},
		v,
	)

	if err != nil {
		return err
	}

	return nil
}

func (steam *Client) getInventory(
	ctx context.Context,
	steamID string,
	v *InventoryResponse,
) error {

	path := steam.config.APIPath + 
	fmt.Sprintf(
		"inventory/%s/%d/%s",
		steamID,
		steam.config.AppID,
		"2")

	err := steam.Client.Get(
		ctx,
		path,
		types.Authorize{},
		v,
	)

	if err != nil {
		return err
	}

	return nil
}