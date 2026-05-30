package steamapis

import (
	"context"

	"github.com/cyrillemad/csmt/types"
)

func (steamapis *Client) HashImageURL(
	hash types.MarketHash) (url string, err error) {
	ctx := context.Background()

	if steamapis.config.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(
			ctx,
			steamapis.config.Timeout,
		)

		defer cancel()
	}

	var response string = ""
	err = steamapis.buildItemImage(hash, &response)
	if err != nil {
		return url, err
	}

	return response, nil
}
