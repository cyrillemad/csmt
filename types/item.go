package types

type Item struct {
	AssetID    string
	AppID      int
	ClassID    string
	InstanceID string
	Amount     int

	Name       string
	MarketName string
	MarketHash MarketHash
	Type       string

	Tradable   bool
	Marketable bool
	Commodity  bool
	IsCurrency bool

	MedianPrice int
	LowestPrice int
	Volume      int

	MarketTradableRestriction   int
	MarketMarketableRestriction int

	BackgroundColor string
	NameColor       string
	IconURL         string
	IconURLLarge    string

	Tags          []ItemTag
	Descriptions  []ItemDescription
	Actions       []ItemAction
	MarketActions []ItemAction
}

type ItemTag struct {
	Category     string
	Name         string
	InternalName string
	Color        string
}

type ItemDescription struct {
	Value string
	Color string
}

type ItemAction struct {
	Name string
	Link string
}
