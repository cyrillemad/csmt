package convert

import (
	"strings"
	"time"
)

func ParseSteamTime(s string) (time.Time, error) {
	parts := strings.SplitN(s, ":", 2)
	return time.ParseInLocation("Jan 02 2006 15", strings.TrimSpace(parts[0]), time.UTC)
}
