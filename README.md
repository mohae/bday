# bday
birthday paradox: calculates the probability of a collision for a number of values within a given range or the number of random values for a given probability of collision

By default, the probability of collision, _p(n;d)_, uses the generalized approximation formula.  A more precise value for __P(A)__, the probability of at least two elements colliding, can be obtained by using the `-c` flag.

The number of items, _d_, that need to be drawn from a set, _n_ (_x^y_), for a given probability, _p_, of at least two colliding is also supported.

The `-d` and `-p` flags are mutually exclusive as the calculations produce one of those values, depending on which flag is present.


## Install


    $ go install github.com/mohae/bday

    // check that it runs; get the flags.
	  $ bday -h
