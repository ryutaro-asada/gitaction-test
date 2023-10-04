package main

import "testing"

func TestAdder(t *testing.T) {

	got := adder(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
