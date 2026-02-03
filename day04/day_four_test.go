package main

import (
	"testing"
)


func TestForkliftMap(t *testing.T) {
	result := ForkliftMap("small_entries.txt")
	resultLarge := ForkliftMap("entries.txt")

	expected := 13
	if result != expected {
		t.Fatalf("ForkliftMap result is %d, want %d", result, expected)
	}

	expectedLarge := 1474
	if resultLarge != expectedLarge {
		t.Fatalf("ForkliftMap result is %d, want %d", resultLarge, expectedLarge)
	}
}

func TestRemovePaperRolls(t *testing.T) {
	result := RemovePaperRolls("small_entries.txt")
	resultLarge := RemovePaperRolls("entries.txt")

	expected := 43
	if result != expected {
		t.Fatalf("RemovePaperRolls result is %d, want %d", result, expected)
	}

	expectedLarge := 8910
	if resultLarge != expectedLarge {
		t.Fatalf("RemovePaperRolls result is %d, want %d", resultLarge, expectedLarge)
	}
}