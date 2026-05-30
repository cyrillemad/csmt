package steamapis

import (
	"fmt"
	"net/url"

	"github.com/cyrillemad/csmt/types"
)

func (steamapis *Client) buildItemImage(
	hash types.MarketHash,
	v *string) error {

	api := fmt.Sprintf(
		"/image/item/%d/%s",
		steamapis.config.AppID,
		url.PathEscape(string(hash)))

	*v = "https://api.steamapis.com" + api

	return nil
}
