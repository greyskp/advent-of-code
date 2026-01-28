package main

import "testing"

func TestIdentifyInvalidIds(t *testing.T) {
	result := IdentifyInvalidIds("small_entries.txt")
	result_large := IdentifyInvalidIds("entries.txt")
	expected := 1227775554
	expected_large := 40398804950
	if result != expected {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result, expected)
	}

	if result_large != expected_large {
		t.Fatalf("IdentifyInvalidIds result for large dataset is %d, want %d", result_large, expected_large)
	}
}
