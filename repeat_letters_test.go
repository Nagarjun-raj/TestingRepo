package main

import "testing"

func TestMain(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"first":  {"interface", "iinntteerrffaaccee"},
		"second": {"pointer", "ppooiinntteerr"},
		"third":  {"channel", "cchhaannnneell"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := doubleChars(tc.input)
			want := tc.want
			if got != want {
				t.Errorf("%s:%s is same as %s", name, got, want)
			}
		})
	}
}
