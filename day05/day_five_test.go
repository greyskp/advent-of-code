package main

import (
	"testing"
)


func TestCountFreshIngredientsFromList(t *testing.T) {
	result := CountFreshIngredientsFromList("small_entries.txt")
	resultLarge := CountFreshIngredientsFromList("entries.txt")

	expected := 3
	if result != expected {
		t.Fatalf("CountFreshIngredientsFromList result is %d, want %d", result, expected)
	}

	expectedLarge := 525
	if resultLarge != expectedLarge {
		t.Fatalf("CountFreshIngredientsFromList result is %d, want %d", resultLarge, expectedLarge)
	}
}

func TestCountAllFreshIngredients(t *testing.T) {
	result := CountAllFreshIngredients("small_entries.txt")
	resultLarge := CountAllFreshIngredients("entries.txt")

	expected := 14
	if result != expected {
		t.Fatalf("CountFreshIngredients result is %d, want %d", result, expected)
	}

	expectedLarge := 333892124923577
	if resultLarge != expectedLarge {
		t.Fatalf("CountFreshIngredients result is %d, want %d", resultLarge, expectedLarge)
	}
}