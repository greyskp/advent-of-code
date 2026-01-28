package main

import "testing"

func TestKnobRotation(t *testing.T) {
	result1, result2 := rotateKnob("small_entries.txt")
	expected1, expected2 := 3, 6
	if result1 != expected1 {
		t.Fatalf("RotateKnob result is %d, want %d", result1, expected1)
	}

	if result2 != expected2 {
		t.Fatalf("RotateKnob part 2 result is %d, want %d", result2, expected2)
	}
}

func TestKnobRotation2(t *testing.T) {
	result1, result2 := rotateKnob("other_entries.txt")
	expected1, expected2 := 5, 22
	if result1 != expected1 {
		t.Fatalf("RotateKnob result is %d, want %d", result1, expected1)
	}

	if result2 != expected2 {
		t.Fatalf("RotateKnob part 2 result is %d, want %d", result2, expected2)
	}
}