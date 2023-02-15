package main

import "testing"

func TestMain(t *testing.T) {
	input := []int{10, 5, 3, 18, 9}
	expected := 18
	got := find(input)
	if expected != got {
		t.Errorf("Got %d but we want %d", got, expected)
	} else {
		t.Logf("Successfully got")
	}

}
