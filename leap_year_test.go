package main

import "testing"

func TestMain(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"first":  {1964, true},
		"second": {2013, false},
		"third":  {1800, false},
		"fourth": {2044, true},
		"fifth":  {2000, true},
		"sixth":  {1991, false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := checkLeapYear(tc.input)
			want := tc.want
			if got != want {
				t.Errorf("%v:%v is not same as %v", name, got, want)
			}
		})
	}

}
