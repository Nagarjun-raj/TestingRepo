package main

import "testing"

func TestPalin(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"first":  {"Malayalam", true},
		"second": {"Nagarjun", false},
		"third":  {"aba", true},
	}

	for name, te := range tests {
		t.Run(name, func(t *testing.T) {
			got := palindromeCheck(te.input)
			want := te.want
			if got != want {
				t.Errorf("%v is not same as %v", got, want)
			} else {
				t.Logf("Successfully got..")
			}
		})
	}

}
