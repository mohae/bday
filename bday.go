package main

import (
	"fmt"
	"math"
)

// d = x^y
// p(n;d) ~ 1-((d-1)/d)^n(n-1)/2
// p(n;d) is the probability that at least two numbers are the same.
func collisionP(n, x, y float64) (p float64, d float64, err error) {
	errPrefix := "can't calculate the probability of at least 2 elements colliding"
	if n <= 0 {
		return 0, 0, fmt.Errorf("%s: n must be > 0", errPrefix)
	}
	if x <= 0 {
		return 0, 0, fmt.Errorf("%s: base must be > 0", errPrefix)
	}
	if y == 0 {
		return 0, 0, fmt.Errorf("%s: if the base is to be used as the upper-end of the range, use 1 as the value of the exponent", errPrefix)
	}
	d = math.Pow(x, y)
	return 1.0 - math.Pow(((d-1.0)/d), (n*(n-1.0))/2.0), d, nil
}

// d = x^y
// n(p;d) ~ SQRT(2d * ln(1/(1-p)))
// n(p;d) is the number of random integers drawn from [1,d] to obtain the
// probility that at least 2 numbers are the same.
func collisionN(p, x, y float64) (n int64, d float64, err error) {
	errPrefix := "can't calculate the number of elements required to obtain the probability that at least 2 elements are the same"
	if p <= 0 {
		return 0, 0, fmt.Errorf("%s: p must be > 0", errPrefix)
	}
	if x <= 0.0 {
		return 0, 0, fmt.Errorf("%s: base must be > 0", errPrefix)
	}
	if y == 0.0 {
		return 0, 0, fmt.Errorf("%s: if the base is to be used as the upper-end of the range, use 1 as the value of the exponent", errPrefix)
	}
	d = math.Pow(x, y)
	// We add 1 because int conversion truncates and if there is a decimal portion
	// it should be the next greater int as minimum n is a whole number.
	return int64(math.Sqrt(2*d*(math.Log((1 / (1 - p))))) + 1.0), d, nil
}

// https://math.dartmouth.edu/archive/m19w03/public_html/Section6-5.pdf
// Theorem 6.15
//In hashing n items into a hash table with k locations, the expected number
// of collisions is: n − k + k(1 − 1/k)^n
func nCollisions(n, x, y float64) (c, d float64, err error) {
	errPrefix := "can't calculate the expected number of collisions when hashing n items into d slots"
	if n <= 0 {
		return 0, 0, fmt.Errorf("%s: n must be > 0", errPrefix)
	}
	if x <= 0 {
		return 0, 0, fmt.Errorf("%s: base must be > 0", errPrefix)
	}
	if y == 0 {
		return 0, 0, fmt.Errorf("%s: if the base is to be used as the number of slots, use 1 as the value of the exponent", errPrefix)
	}
	d = math.Pow(x, y)
	return n - d + (d * math.Pow(1-(1.0/d), n)), d, nil
}
