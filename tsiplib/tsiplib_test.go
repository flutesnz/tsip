package tsiplib

import "testing"

func TestHello(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"world", "Welcome, world. "},
		{"dev", "Welcome, dev. "},
	}

	for _, c := range cases {
		got := Hello(c.in)
		if got != c.want {
			t.Errorf("Hello(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
