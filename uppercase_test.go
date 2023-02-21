package main

import "testing"

func TestMain(t *testing.T) {
	name := "praShAnth"
	actual := ToUppercase(name)
	expected := "PRASHANTH"
	if actual != expected {
		t.Errorf("%s is not same as %s", actual, expected)
	}
}
