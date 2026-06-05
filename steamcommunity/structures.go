package steamcommunity

type PriceOverviewResponse struct {
	Status      bool   `json:"success"`
	LowestPrice string `json:"lowest_price"`
	MedianPrice string `json:"median_price"`
	Volume      string `json:"volume"`
}

type RenderSearchResponse struct {
	Status     bool `json:"success"`
	Start      int  `json:"start"`
	PageSize   int  `json:"pagesize"`
	TotalCount int  `json:"total_count"`

	SearchData SearchData `json:"searchdata"`

	Results []SearchResult `json:"results"`
}

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

type SearchData struct {
	Query              string `json:"query"`
	SearchDescriptions bool   `json:"search_descriptions"`

	TotalCount int `json:"total_count"`
	PageSize   int `json:"pagesize"`

	Prefix      string `json:"prefix"`
	ClassPrefix string `json:"class_prefix"`
}

type SearchResult struct {
	Name     string `json:"name"`
	HashName string `json:"hash_name"`

	SellListings int `json:"sell_listings"`

	SellPrice     int    `json:"sell_price"`
	SellPriceText string `json:"sell_price_text"`

	AppIcon string `json:"app_icon"`
	AppName string `json:"app_name"`

	AssetDescription AssetDescription `json:"asset_description"`

	SalePriceText string `json:"sale_price_text"`
}

type AssetDescription struct {
	AppID int `json:"appid"`

	ClassID    string `json:"classid"`
	InstanceID string `json:"instanceid"`

	BackgroundColor string `json:"background_color"`

	IconURL string `json:"icon_url"`

	Tradable int `json:"tradable"`

	Name string `json:"name"`

	NameColor string `json:"name_color"`

	Type string `json:"type"`

	MarketName     string `json:"market_name"`
	MarketHashName string `json:"market_hash_name"`

	Commodity int `json:"commodity"`
}
type SortColumn string

const (
	SortColumnName     SortColumn = "name"
	SortColumnPrice    SortColumn = "price"
	SortColumnQuantity SortColumn = "quantity"
	SortColumnPopular  SortColumn = "popular"
)

type SortDirection string

const (
	SortAsc  SortDirection = "asc"
	SortDesc SortDirection = "desc"
)

type RenderSearchOptions struct {
	Query              string
	Start              int
	Count              int
	SortColumn         SortColumn
	SortDirection      SortDirection
	SearchDescriptions bool
	PriceMin           int
	PriceMax           int
	NoRender           bool
}
