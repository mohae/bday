# bday
Performs estimates related to hashing collisions and probabilities.  

For estimating the probability of collision for a given number of elements being inserted into an array with _x^y_ elements, the birthday paradox is used.

## Calculations
### Probability of collision for n elements.
`-n` estimates the probability of at least two elements colliding, _p(n;d)_, when inserting _n_ elements into an array, _d_.  

    $ bday -n 10 -x 365
	$ The probability of at least two elements colliding when hashing 10 elements into a 365 element array is 11.614024%.

### Number of elements for a given probability of collision.
`-p` estimates the number of elements, _n_, that need to be drawn from a set, _d_ (_x^y_), for a collision to occur given a probability, _p_.

    $ bday -p 11.6 -x 365
	$ 10 elements are needed to have a of 11.6000% chance of 2 elements colliding when hashing into a 365 element array.

### Number of collisions when inserting n elements.
Use the `-c` flag in conjunction with the `-n` flag to calculate the expected number of collisions when hashing _n_ elements into _d_ slots.  This uses the __Theorem 6.15__ from https://math.dartmouth.edu/archive/m19w03/public_html/Section6-5.pdf.

    $ bday -n 100 -x 10 -y3 -c
    $ The estimated number of collisions when hashing 100 elements into 256 slots is 17.08581506804407.

### Flag notes
The `-n` and `-p` flags are mutually exclusive as the calculations produce one of those values, depending on which flag is present.

If the base, `-x`, is to be used as the size of the array, don't use the exponent flag, `-y`.

## Install

    $ go install github.com/mohae/bday

    // check that it runs; get the flags.
    $ bday -h
