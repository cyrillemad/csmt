package convert

import (
	"testing"
	"time"
)

func TestParseSteamTime(t *testing.T) {
	got, err := ParseSteamTime("Nov 01 2013 01: +0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := time.Date(2013, time.November, 1, 1, 0, 0, 0, time.UTC)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
