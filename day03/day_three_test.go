package main

import "testing"

func TestHighestJoltage(t *testing.T) {
	result := HighestJoltage("small_entries.txt")
	result_large := HighestJoltage("entries.txt")
	expected := 357
	if result != expected {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result, expected)
	}

	expected_large := 17408
	if result != expected {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result_large, expected_large)
	}
}
