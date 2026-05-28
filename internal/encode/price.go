package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParsePrice(value string) (float64, error) {
	re := regexp.MustCompile(`[0-9.,]+`)

	match := re.FindString(value)

	if match == "" {
		return 0, fmt.Errorf("invalid price: %s", value)
	}

	match = strings.ReplaceAll(match, ",", ".")

	return strconv.ParseFloat(match, 64)
}
