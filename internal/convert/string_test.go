package convert

import "testing"

func TestInt(t *testing.T) {
	value, err := Int("42")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if value != 42 {
		t.Fatalf("unexpected value: %d", value)
	}

	_, err = Int("not-a-number")
	if err == nil {
		t.Fatal("expected error for invalid input")
	}
}
