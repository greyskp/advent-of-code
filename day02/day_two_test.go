package main

import "testing"

func TestIdentifyInvalidIds(t *testing.T) {
	result := IdentifyInvalidIds("small_entries.txt")
	expected := 1227775554
	if result != expected {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result, expected)
	}
}
