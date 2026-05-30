package csmt

import (
	"github.com/cyrillemad/csmt/internal/httpclient"
	"github.com/cyrillemad/csmt/steamapis"
	steam "github.com/cyrillemad/csmt/steamcommunity"
)

type NoAuthorizeClient struct {
	Community steam.Client
	Apis      steamapis.Client
}

func NewNoAuthClient(
	options ...any) *NoAuthorizeClient {

	var communityOptions []steam.Option
	var apiOptions []steamapis.Option
	var httpOptions []httpclient.Option

	for _, option := range options {
		switch o := option.(type) {

		case steam.Option:
			communityOptions = append(
				communityOptions,
				o,
			)

		case steamapis.Option:
			apiOptions = append(
				apiOptions,
				o,
			)

		case httpclient.Option:
			httpOptions = append(
				httpOptions,
				o,
			)
		}
	}

	if len(httpOptions) > 0 {
		for _, option := range httpOptions {
			apiOptions = append(
				apiOptions,
				steamapis.WithHTTPOption(option))
			communityOptions = append(
				communityOptions,
				steam.WithHTTPOption(option))
		}
	}
	steamCommunity := steam.NewClient(communityOptions...)
	steamApis := steamapis.NewClient(apiOptions...)

	return &NoAuthorizeClient{
		Community: *steamCommunity,
		Apis:      *steamApis,
	}
}
