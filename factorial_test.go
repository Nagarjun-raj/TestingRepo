package main

import "testing"

func TestFactorial(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"first":  {6, 720},
		"second": {5, 120},
		"third":  {7, 5040},
	}
	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			got := factorialCalc(te.input)
			if te.want != got {
				t.Fatalf("we got %d but we wanted %d", got, te.want)
			} else {
				t.Logf("Successfully got..")
			}
		})
	}

}
