package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	input := []int{12, 34, 90, 27, 65}
	expected := []int{90, 65}
	actual := search(input)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("Successfull")
		return
	}
	t.Errorf("%v is not same as %v", expected, actual)
}
