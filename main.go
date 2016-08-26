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
	x float64 // base of the exponent
	y float64 // exponent
	c bool
	h bool
)

func init() {
	flag.Float64Var(&n, "n", 0, "the number of random values to use for collision probability calculations")
	flag.Float64Var(&p, "p", 0, "the probability for which collision calculation will be made, as a percentage (without the %% symbol)")
	flag.Float64Var(&x, "x", 0, "the base of the exponent")
	flag.Float64Var(&y, "y", 1, "the exponent; do not use if the base is to be used as the upper end of the range of values")
	flag.BoolVar(&c, "c", false, "calculate the estimated number of collisions for n items in x^y slots; can only be used with -n")
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
	if int64(p) <= 0 && n <= 0 {
		fmt.Fprintf(os.Stderr, "either -p or -n must be specified and > 0\n")
		return printHelpSuggestion()
	}
	if p > 0 && n > 0 {
		fmt.Fprintf(os.Stderr, "-p and -n are mutually exclusive, only one can be specified\n")
		return printHelpSuggestion()
	}
	if x <= 0 {
		fmt.Fprintf(os.Stderr, "the base, -x, must be > 0\n")
		return printHelpSuggestion()
	}
	// even though y defaults to 1; make sure it wasn't set to an invalid number.
	if y <= 0 {
		fmt.Fprintf(os.Stderr, "the exponent, -y, must be >= 1\n")
		return printHelpSuggestion()
	}
	if n <= 0 && c {
		fmt.Fprintf(os.Stderr, "collision estimates can only be done when -n is provided and > 0\n")
		return printHelpSuggestion()
	}
	if n > 0 {
		if c {
			c, d, err := nCollisions(n, x, y)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error estimating number of collisions: %s\n", err)
				return 1
			}
			fmt.Printf("The estimated number of collisions when hashing %g elements into %g slots is %g.\n", n, d, c)
			return 0
		}
		// calculate probability that at least 2 values are the same
		prob, d, err := collisionP(n, x, y)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error calculating collision probability: %s\n", err)
			return 1
		}
		fmt.Printf("The probability at least two elements colliding when hashing %g elements into a %g element array is %f%%.\n", n, d, prob*100.0)
		return 0
	}
	// calculate the number of elements needed for a given collision probability
	n, d, err := collisionN(p/100.0, x, y)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calculating number of elements needed for a given collision probability: %s\n", err)
		return 1
	}
	fmt.Printf("%g elements are needed to have a %f%% chance of 2 elements colliding when hashing into a %g element array.\n", float64(n), p, d)
	return 0
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  Either -p or -n must be specified; these are mutually exclusive.\n")
	flag.PrintDefaults()
}

func printHelpSuggestion() int {
	fmt.Fprintf(os.Stderr, "use -h for help and usage output]n")
	return 1
}
