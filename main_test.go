package main

import "testing"

func TestExample(t *testing.T) {
	got := 2 + 2
	want := 4

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
