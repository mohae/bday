package main

import (
	"testing"
)

func TestCollisionP(t *testing.T) {
	tests := []struct {
		n, x, y float64
		p, d    float64
		err     string
	}{
		{0, 0, 0, 0, 0, "can't calculate the probability of at least 2 elements colliding: n must be > 0"},
		{1, 0, 0, 0, 0, "can't calculate the probability of at least 2 elements colliding: base must be > 0"},
		{1, 10, 0, 0, 0, "can't calculate the probability of at least 2 elements colliding: if the base is to be used as the upper-end of the range, use 1 as the value of the exponent"},
		{10, 10, 6, 0, 1000000, ""},
		{2, 10, 1, 9.9, 10, ""},

		{100, 10, 6, 0.5, 1000000, ""},
		{1000, 10, 6, 39.3, 1000000, ""},
		{1, 365, 1, 0, 365, ""},
		{5, 365, 1, 2.7, 365, ""},
		{10, 365, 1, 11.6, 365, ""},

		{60, 365, 1, 99.2, 365, ""},
		{70, 365, 1, 99.9, 365, ""},
		{1000000, 64, 8, 0.2, 281474976710656, ""},
	}
	for i, test := range tests {
		p, d, err := collisionP(test.n, test.x, test.y)
		if err != nil {
			if err.Error() != test.err {
				t.Errorf("%d error: got %q want %q", i, err.Error(), test.err)
			}
			continue
		}
		if test.err != "" {
			t.Errorf("%d: got no error, want %q", i, test.err)
		}
		// we turn it into an int to simplify comparison.
		if int(p*100.0) != int(test.p) {
			t.Errorf("%d p: got %3.1f; want %3.1f", i, p*100.0, test.p)
		}
		if d != test.d {
			t.Errorf("%d d: got %f; want %f", i, d, test.d)
		}
	}
}

func TestCollisionN(t *testing.T) {
	tests := []struct {
		p, x, y float64
		n       int64
		d       float64
		err     string
	}{
		{0, 0, 0, 0, 0, "can't calculate the number of elements required to obtain the probability that at least 2 elements are the same: p must be > 0"},
		{0, 1, 0, 0, 0, "can't calculate the number of elements required to obtain the probability that at least 2 elements are the same: p must be > 0"},
		{.1, 0, 0, 0, 0, "can't calculate the number of elements required to obtain the probability that at least 2 elements are the same: base must be > 0"},
		{.1, 1, 0, 0, 0, "can't calculate the number of elements required to obtain the probability that at least 2 elements are the same: if the base is to be used as the upper-end of the range, use 1 as the value of the exponent"},
		{.1, 10, 4, 46, 10000, ""},

		{.1, 10, 2, 5, 100, ""},
		{.1, 10, 3, 15, 1000, ""},
		{.5, 10, 3, 38, 1000, ""},
		{.117, 365, 1, 10, 365, ""},
		{.5, 365, 1, 23, 365, ""},
	}
	for i, test := range tests {
		n, d, err := collisionN(test.p, test.x, test.y)
		if err != nil {
			if err.Error() != test.err {
				t.Errorf("%d: got %q; want %q", i, err, test.err)
			}
			continue
		}
		if test.err != "" {
			t.Errorf("%d: got no error; want %q; ", i, test.err)
			continue
		}
		if int(d) != int(test.d) {
			t.Errorf("%d d: got %d; want %d", i, int(d), int(test.d))
		}
		if n != test.n {
			t.Errorf("%d n: got %d; want %d", i, n, test.n)
		}
	}
}
