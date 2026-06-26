package encode

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cyrillemad/csmt/internal/convert"
)

type PriceHistoryEntry struct {
	Time   time.Time
	Price  float64
	Volume int
}

func ParsePriceHistory(prices [][]any) ([]PriceHistoryEntry, error) {
	entries := make([]PriceHistoryEntry, 0, len(prices))
	for _, row := range prices {
		if len(row) != 3 {
			return nil, fmt.Errorf("unexpected row length %d", len(row))
		}

		dateStr, err := anyToString(row[0])
		if err != nil {
			return nil, fmt.Errorf("parse date: %w", err)
		}

		t, err := convert.ParseSteamTime(dateStr)
		if err != nil {
			return nil, fmt.Errorf("parse time %q: %w", dateStr, err)
		}

		priceStr, err := anyToString(row[1])
		if err != nil {
			return nil, fmt.Errorf("parse price: %w", err)
		}

		price, err := ParsePrice(priceStr)
		if err != nil {
			return nil, fmt.Errorf("parse price %q: %w", priceStr, err)
		}

		volStr, err := anyToString(row[2])
		if err != nil {
			return nil, fmt.Errorf("parse volume: %w", err)
		}

		volume, err := strconv.Atoi(strings.ReplaceAll(volStr, ",", ""))
		if err != nil {
			return nil, fmt.Errorf("parse volume %q: %w", volStr, err)
		}

		entries = append(entries, PriceHistoryEntry{
			Time:   t,
			Price:  price,
			Volume: volume,
		})
	}
	return entries, nil
}

func anyToString(v any) (string, error) {
	switch x := v.(type) {
	case string:
		return x, nil
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("unexpected type %T", v)
	}
}
