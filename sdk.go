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

	httpClient := httpclient.NewClient(httpOptions...)
	steamCommunity := steam.NewClient(httpClient, communityOptions...)
	steamApis := steamapis.NewClient(httpClient, apiOptions...)

	return &NoAuthorizeClient{
		Community: *steamCommunity,
		Apis:      *steamApis,
	}
}
