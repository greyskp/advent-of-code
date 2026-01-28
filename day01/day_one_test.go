package main

import "testing"

func TestKnobRotation(t *testing.T) {
	expected1, expected2 := rotateKnob("small_entries.txt")
	want1, want2 := 3, 6
	if expected1 != want1 {
		t.Fatalf("RotateKnob result is %d, want %d", expected1, want1)
	}

	if expected2 != want2 {
		t.Fatalf("RotateKnob part 2 result is %d, want %d", expected2, want2)
	}
}

func TestKnobRotation2(t *testing.T) {
	expected1, expected2 := rotateKnob("other_entries.txt")
	want1, want2 := 5, 22
	if expected1 != want1 {
		t.Fatalf("RotateKnob result is %d, want %d", expected1, want1)
	}

	if expected2 != want2 {
		t.Fatalf("RotateKnob part 2 result is %d, want %d", expected2, want2)
	}
}