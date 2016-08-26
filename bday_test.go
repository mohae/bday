package main

import "testing"

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

func TestNCollisions(t *testing.T) {
	tests := []struct {
		n, x, y, c, d float64
		err           string
	}{
		{0, 0, 0, 0, 0, "can't calculate the expected number of collisions when hashing n items into d slots: n must be > 0"},
		{0, 1, 0, 0, 0, "can't calculate the expected number of collisions when hashing n items into d slots: n must be > 0"},
		{1, 0, 0, 0, 0, "can't calculate the expected number of collisions when hashing n items into d slots: base must be > 0"},
		{1, 1, 0, 0, 0, "can't calculate the expected number of collisions when hashing n items into d slots: if the base is to be used as the number of slots, use 1 as the value of the exponent"},
		{1, 10, 1, 0, 10, ""},

		{2, 10, 1, .1, 10, ""},
		{5, 10, 1, .9049, 10, ""},
		{1, 10, 2, 0, 100, ""},
		{2, 10, 2, .0099, 100, ""},
		{10, 10, 2, .4382, 100, ""},

		{20, 10, 2, 1.7906, 100, ""},
		{25, 10, 2, 2.7821, 100, ""},
		{10, 10, 3, .0448, 1000, ""},
		{20, 10, 3, .1888, 1000, ""},
		{50, 10, 3, 1.2056, 1000, ""},

		{100, 10, 3, 4.7921, 1000, ""},
		{200, 10, 3, 18.6488, 1000, ""},
		{300, 10, 3, 40.7070, 1000, ""},
		{500, 10, 3, 106.3789, 1000, ""},
		{800, 10, 3, 249.1491, 1000, ""},
	}
	for i, test := range tests {
		c, d, err := nCollisions(test.n, test.x, test.y)
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
		// just check up to ten thousandths
		if int(c*1000.0) != int(test.c*1000.0) {
			t.Errorf("%d: %g items in %g slots:: got %g; want %g", i, test.n, d, c, test.c)
		}
	}
}
