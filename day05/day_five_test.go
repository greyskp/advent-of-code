package main

import (
	"testing"
)


func TestCountFreshIngredients(t *testing.T) {
	result := CountFreshIngredients("small_entries.txt")
	resultLarge := CountFreshIngredients("entries.txt")

	expected := 3
	if result != expected {
		t.Fatalf("CountFreshIngredients result is %d, want %d", result, expected)
	}

	expectedLarge := 525
	if resultLarge != expectedLarge {
		t.Fatalf("CountFreshIngredients result is %d, want %d", resultLarge, expectedLarge)
	}
}