package main

import (
	"testing"
)

func TestHighestJoltagePartTwo(t *testing.T) {
	result := HighestJoltagePartTwo("small_entries.txt", 2)
	result3 := HighestJoltagePartTwo("small_entries.txt", 3)
	result4 := HighestJoltagePartTwo("small_entries.txt", 4)
	result12 := HighestJoltagePartTwo("small_entries.txt", 12)
	resultLarge := HighestJoltagePartTwo("entries.txt", 12)

	expected := 357
	if result != expected {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result, expected)
	}

	expected3 := 3205
	if result3 != expected3 {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result3, expected3)
	}

	expected4 := 31684
	if result4 != expected4 {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result4, expected4)
	}

	expected12 := 3121910778619
	if result12 != expected12 {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", result12, expected12)
	}

	expectedLarge := 172740584266849
	if resultLarge != expectedLarge {
		t.Fatalf("IdentifyInvalidIds result is %d, want %d", resultLarge, expectedLarge)
	}
}
