package encode

import (
	"encoding/json"
	"testing"

	"github.com/cyrillemad/csmt/types"
)

const sampleInventory = `{
	"assets": [
		{
			"appid": 730,
			"contextid": "2",
			"assetid": "123",
			"classid": "1",
			"instanceid": "0",
			"amount": "1"
		}
	],
	"descriptions": [
		{
			"appid": 730,
			"classid": "1",
			"instanceid": "0",
			"currency": 0,
			"background_color": "393b3e",
			"icon_url": "icon",
			"icon_url_large": "icon_large",
			"tradable": 1,
			"name": "AK-47 | Redline",
			"name_color": "D2D2D2",
			"type": "Classified Rifle",
			"market_name": "AK-47 | Redline (Field-Tested)",
			"market_hash_name": "AK-47 | Redline (Field-Tested)",
			"commodity": 0,
			"market_tradable_restriction": 7,
			"market_marketable_restriction": 0,
			"marketable": 1,
			"tags": [
				{
					"category": "Type",
					"internal_name": "CSGO_Type_Rifle",
					"localized_category_name": "Type",
					"localized_tag_name": "Rifle"
				}
			],
			"descriptions": [
				{
					"value": "Exterior: Field-Tested",
					"color": "9da1a9"
				}
			],
			"actions": [
				{
					"link": "steam://inspect/730/2/123",
					"name": "Inspect in Game..."
				}
			]
		}
	],
	"total_inventory_count": 1,
	"success": 1
}`

func TestParseInventoryItems(t *testing.T) {
	raw := InventoryResponse{}
	if err := json.Unmarshal([]byte(sampleInventory), &raw); err != nil {
		t.Fatalf("unmarshal fixture: %v", err)
	}

	items, err := ParseInventoryItems(raw)
	if err != nil {
		t.Fatalf("ParseInventoryItems: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	item := items[0]
	if item.AssetID != "123" {
		t.Fatalf("unexpected asset id: %q", item.AssetID)
	}
	if item.Name != "AK-47 | Redline" {
		t.Fatalf("unexpected name: %q", item.Name)
	}
	if item.MarketHash != types.MarketHash("AK-47 | Redline (Field-Tested)") {
		t.Fatalf("unexpected market hash: %q", item.MarketHash)
	}
	if !item.Tradable || !item.Marketable {
		t.Fatal("expected tradable and marketable item")
	}
	if item.Commodity || item.IsCurrency {
		t.Fatal("expected regular item, not commodity or currency")
	}
	if item.Amount != 1 {
		t.Fatalf("unexpected amount: %d", item.Amount)
	}
	if len(item.Tags) != 1 || item.Tags[0].Name != "Rifle" {
		t.Fatalf("unexpected tags: %+v", item.Tags)
	}
	if len(item.Descriptions) != 1 ||
		item.Descriptions[0].Value != "Exterior: Field-Tested" {
		t.Fatalf("unexpected descriptions: %+v", item.Descriptions)
	}
}
