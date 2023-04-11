package main

import "testing"

func TestMain(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"first":  {"Globallogic", "cigollabolG"},
		"second": {"Bangalore", "erolagnaB"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := reverse(tc.input)
			want := tc.want
			if got != want {
				t.Errorf("%v:%s is not same as %s", name, got, want)
			}
		})
	}
}
