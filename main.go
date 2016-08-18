package main

import (
	"flag"
	"fmt"
	"os"
)

// if an exponent isn't passed, the base will be used as the upper value of
// the range
var (
	p float64 // probability
	n float64 // number of random values
	x int64   // base of the exponent
	y int64   // exponent
	h bool
)

func init() {
	flag.Float64Var(&n, "n", 0, "the number of random values to use for collision probability calculations")
	flag.Float64Var(&p, "p", 0, "the probability for which collision calculation will be made, as a percentage (without the %% symbol)")
	flag.Int64Var(&x, "x", 0, "the base of the exponent")
	flag.Int64Var(&y, "y", 1, "the exponent; do not use if the base is to be used as the upper end of the range of values")
	flag.BoolVar(&h, "h", false, "help output")
	flag.BoolVar(&h, "help", false, "help output")
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()
	if h {
		usage()
		return 1
	}
	// either the probability or the number of random values must be passed.
	if int64(p) == 0 && n == 0 {
		fmt.Fprintf(os.Stderr, "either -p or -n must be specified\n")
		fmt.Fprintf(os.Stderr, "use -h for help and usage output\n")
		return 1
	}
	if p > 0 && n != 0 {
		fmt.Fprintf(os.Stderr, "-p and -n are mutually exclusive, only one can be specified\n")
		fmt.Fprintf(os.Stderr, "use -h for help and usage output\n")
		return 1
	}
	if x == 0 {
		fmt.Fprintf(os.Stderr, "the base, -x, must be > 0\n")
		fmt.Fprintf(os.Stderr, "use -h for help and usage output]n")
		return 1
	}

	if n > 0 {
		// calculate probability that at least 2 values are the same
		prob, d, err := collisionP(float64(n), float64(x), float64(y))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error calculating collision probability: %s\n", err)
			return 1
		}
		fmt.Printf("The probability that %g elements for %g slots collide is: %f\n", n, d, prob*100.0)
		return 0
	}
	// calculate the number of elements needed for a given collision probability
	n, d, err := collisionN(p/100.0, float64(x), float64(y))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calculating number of elements needed for a given collision probability: %s\n", err)
		return 1
	}
	fmt.Printf("%g elements are needed to have a %f probability of collision for %g slots\n", float64(n), p/100.0, d)
	return 0
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  Either -p or -n must be specified; these are mutually exclusive.\n")
	flag.PrintDefaults()
}
