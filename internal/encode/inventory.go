package encode

import (
	"fmt"

	"github.com/cyrillemad/csmt/internal/convert"
	"github.com/cyrillemad/csmt/types"
)

type InventoryResponse struct {
	Assets []InventoryAsset `json:"assets"`

	Descriptions []InventoryDescription `json:"descriptions"`

	TotalInventoryCount int `json:"total_inventory_count"`

	Success int `json:"success"`

	MoreItems bool   `json:"more_items,omitempty"`
	LastAsset string `json:"last_assetid,omitempty"`

	RWGRSN int `json:"rwgrsn,omitempty"`
}

type InventoryAsset struct {
	AppID int `json:"appid"`

	ContextID string `json:"contextid"`

	AssetID string `json:"assetid"`

	ClassID string `json:"classid"`

	InstanceID string `json:"instanceid"`

	Amount string `json:"amount"`
}

type InventoryDescription struct {
	AppID int `json:"appid"`

	ClassID string `json:"classid"`

	InstanceID string `json:"instanceid"`

	Currency int `json:"currency"`

	BackgroundColor string `json:"background_color"`

	IconURL      string `json:"icon_url"`
	IconURLLarge string `json:"icon_url_large"`

	Descriptions []ItemDescription `json:"descriptions"`

	Tradable int `json:"tradable"`

	Actions []InventoryAction `json:"actions"`

	Name string `json:"name"`

	NameColor string `json:"name_color"`

	Type string `json:"type"`

	MarketName string `json:"market_name"`

	MarketHashName string `json:"market_hash_name"`

	Commodity int `json:"commodity"`

	MarketTradableRestriction int `json:"market_tradable_restriction"`

	MarketMarketableRestriction int `json:"market_marketable_restriction"`

	Marketable int `json:"marketable"`

	Tags []InventoryTag `json:"tags"`

	MarketActions []InventoryAction `json:"market_actions"`
}

type ItemDescription struct {
	Value string `json:"value"`

	Color string `json:"color,omitempty"`
}

type InventoryAction struct {
	Link string `json:"link"`

	Name string `json:"name"`
}

type InventoryTag struct {
	Category string `json:"category"`

	InternalName string `json:"internal_name"`

	LocalizedCategoryName string `json:"localized_category_name"`

	LocalizedTagName string `json:"localized_tag_name"`

	Color string `json:"color,omitempty"`
}

func ParseInventoryItems(raw InventoryResponse) ([]types.Item, error) {
	if !convert.Success(raw.Success) {
		return nil, fmt.Errorf("steam returned non-success inventory response")
	}

	descriptions := make(
		map[string]InventoryDescription,
		len(raw.Descriptions),
	)
	for _, description := range raw.Descriptions {
		key := convert.InventoryKey(
			description.AppID,
			description.ClassID,
			description.InstanceID,
		)
		descriptions[key] = description
	}

	items := make([]types.Item, 0, len(raw.Assets))
	for _, asset := range raw.Assets {
		key := convert.InventoryKey(asset.AppID, asset.ClassID, asset.InstanceID)
		description, ok := descriptions[key]
		if !ok {
			continue
		}

		amount, err := convert.Int(asset.Amount)
		if err != nil {
			return nil, fmt.Errorf(
				"parse amount for asset %s: %w",
				asset.AssetID,
				err,
			)
		}

		items = append(items, itemFromInventory(asset, description, amount))
	}

	return items, nil
}

func itemFromInventory(
	asset InventoryAsset,
	description InventoryDescription,
	amount int,
) types.Item {
	tags := convert.Map(description.Tags, func(tag InventoryTag) types.ItemTag {
		return types.ItemTag{
			Category:     tag.LocalizedCategoryName,
			Name:         tag.LocalizedTagName,
			InternalName: tag.InternalName,
			Color:        tag.Color,
		}
	})

	descriptions := convert.Map(
		description.Descriptions,
		func(line ItemDescription) types.ItemDescription {
			return types.ItemDescription{
				Value: line.Value,
				Color: line.Color,
			}
		},
	)

	actions := convert.Map(description.Actions, func(action InventoryAction) types.ItemAction {
		return types.ItemAction{
			Name: action.Name,
			Link: action.Link,
		}
	})

	marketActions := convert.Map(
		description.MarketActions,
		func(action InventoryAction) types.ItemAction {
			return types.ItemAction{
				Name: action.Name,
				Link: action.Link,
			}
		},
	)

	return types.Item{
		AssetID:    asset.AssetID,
		AppID:      asset.AppID,
		ClassID:    asset.ClassID,
		InstanceID: asset.InstanceID,
		Amount:     amount,

		Name:       description.Name,
		MarketName: description.MarketName,
		MarketHash: types.MarketHash(description.MarketHashName),
		Type:       description.Type,

		Tradable:   convert.Flag(description.Tradable),
		Marketable: convert.Flag(description.Marketable),
		Commodity:  convert.Flag(description.Commodity),
		IsCurrency: convert.Flag(description.Currency),

		MarketTradableRestriction:   description.MarketTradableRestriction,
		MarketMarketableRestriction: description.MarketMarketableRestriction,

		BackgroundColor: description.BackgroundColor,
		NameColor:       description.NameColor,
		IconURL:         description.IconURL,
		IconURLLarge:    description.IconURLLarge,

		Tags:          tags,
		Descriptions:  descriptions,
		Actions:       actions,
		MarketActions: marketActions,
	}
}
