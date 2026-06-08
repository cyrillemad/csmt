package convert

import "testing"

func TestFlag(t *testing.T) {
	if Flag(0) {
		t.Fatal("expected false for 0")
	}
	if !Flag(1) {
		t.Fatal("expected true for 1")
	}
	if Flag(2) {
		t.Fatal("expected false for values other than 1")
	}
}

func TestSuccess(t *testing.T) {
	if !Success(1) {
		t.Fatal("expected success for 1")
	}
	if Success(0) {
		t.Fatal("expected failure for 0")
	}
}

func TestInventoryKey(t *testing.T) {
	key := InventoryKey(730, "1", "0")
	if key != "730_1_0" {
		t.Fatalf("unexpected key: %q", key)
	}
}
