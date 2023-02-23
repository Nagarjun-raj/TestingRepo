package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	names := []string{"Roy", "Ted", "Nick", "Chris", "Tony"}
	expected := map[string]int{"C": 1, "N": 1, "R": 1, "T": 2}
	actual := countFirstLetter(names)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("Successfull")
		return
	}
	t.Errorf("%v is not same as %v", expected, actual)
}
