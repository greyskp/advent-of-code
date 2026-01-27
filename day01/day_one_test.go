package main

import "testing"

func TestKnobRotation(t *testing.T) {
	got := rotateKnob("small_entries.txt")
	want := 3
	if got != want {
		t.Fatalf("RotateKnob result is %d, want %d", got, want)
	}
}
