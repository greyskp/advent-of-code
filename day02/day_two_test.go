package main

import "testing"

func TestIdentifyInvalidIds(t *testing.T) {
	result := IdentifyInvalidIds("small_entries.txt")
	resultLarge := IdentifyInvalidIds("entries.txt")
	expected := 1227775554
	expectedLarge := 40398804950
	if result != expected {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result, expected)
	}

	if resultLarge != expectedLarge {
		t.Fatalf("IdentifyInvalidIds result for large dataset is %d, want %d", resultLarge, expectedLarge)
	}
}

func TestIdentifyInvalidIdsPartTwo(t *testing.T) {
	result := InvalidIdsPartTwo("small_entries.txt")
	expected := 4174379265
	if result != expected {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result, expected)
	}
}
