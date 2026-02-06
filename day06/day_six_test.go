package main

import (
	"testing"
)


func TestMathSolver(t *testing.T) {
	result := MathSolver("small_entries.txt")
	resultLarge := MathSolver("entries.txt")

	expected := 4277556
	if result != expected {
		t.Fatalf("MathSolver result is %d, want %d", result, expected)
	}

	expectedLarge := 6343365546996
	if resultLarge != expectedLarge {
		t.Fatalf("MathSolver result is %d, want %d", resultLarge, expectedLarge)
	}
}
