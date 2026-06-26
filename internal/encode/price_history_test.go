package encode

import (
	"testing"
	"time"
)

func TestParsePriceHistory(t *testing.T) {
	prices := [][]any{
		{"Nov 01 2013 01: +0", "4.50", "1"},
		{"Dec 15 2020 12: +0", 126.34, "42"},
	}

	entries, err := ParsePriceHistory(prices)
	if err != nil {
		t.Fatalf("ParsePriceHistory: %v", err)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	first := entries[0]
	if first.Price != 4.50 {
		t.Errorf("expected price 4.50, got %v", first.Price)
	}
	if first.Volume != 1 {
		t.Errorf("expected volume 1, got %d", first.Volume)
	}
	want := time.Date(2013, time.November, 1, 1, 0, 0, 0, time.UTC)
	if first.Time != want {
		t.Errorf("expected time %v, got %v", want, first.Time)
	}

	second := entries[1]
	if second.Price != 126.34 {
		t.Errorf("expected price 126.34, got %v", second.Price)
	}
	if second.Volume != 42 {
		t.Errorf("expected volume 42, got %d", second.Volume)
	}
}
