package csmt

import (
	"github.com/cyrillemad/csmt/steamapis"
	steam "github.com/cyrillemad/csmt/steamcommunity"
)

type NoAuthorizeClient struct {
	Community steam.Client
	Apis      steamapis.Client
}

func NewNoAuthClient() *NoAuthorizeClient {

	steamCommunity := steam.NewClient()
	steamApis := steamapis.NewClient()

	return &NoAuthorizeClient{
		Community: *steamCommunity,
		Apis:      *steamApis,
	}
}
