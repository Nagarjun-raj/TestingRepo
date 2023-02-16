package main

import (
	"reflect"
	"testing"
)

func SortTest(t *testing.T) {
	arr := []int{5, 2, 8, 1, 15}
	want := []int{1, 2, 5, 8, 15}
	got := sorting(arr)
	if reflect.DeepEqual(want, got) {
		t.Logf("Successfully got the output")
	} else {
		t.Errorf("We got %d but we want %d", got, want)
	}

}
