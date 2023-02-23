package main

import (
	"reflect"
	"testing"
)

func TestStringSlice(t *testing.T) {
	input := []string{"nagarjun", "arjun", "Raju", "kumar"}
	want := []string{"NagarjuN", "ArjuN", "RajU", "KumaR"}
	got := capiFirstLast(input)
	if reflect.DeepEqual(want, got) {
		t.Logf("Successful")
	} else {
		t.Error("Want is not same as got")
	}

}
