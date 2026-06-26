package steamcommunity

import (
	"strings"

	"github.com/cyrillemad/csmt/types"
)

type Cookie struct {
	SteamLoginSecure string
	SessionID        string
}

func (c Cookie) header() types.Cookie {
	parts := make([]string, 0, 2)
	if c.SteamLoginSecure != "" {
		parts = append(parts, "steamLoginSecure="+c.SteamLoginSecure)
	}
	if c.SessionID != "" {
		parts = append(parts, "sessionid="+c.SessionID)
	}
	result := strings.Join(parts, "; ")
	return types.Cookie(result)
}

func (c Cookie) empty() bool {
	return c.SteamLoginSecure == "" && c.SessionID == ""
}
