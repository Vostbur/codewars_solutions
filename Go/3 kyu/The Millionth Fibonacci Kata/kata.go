package kata

import "math/big"

func FibDinamic(n int64) *big.Int {
	// Codewars returns "Execution Timed Out (12000 ms)" on fib(2000000)
	var a, b = big.NewInt(0), big.NewInt(1)
	var neg bool

	if n < 0 {
		n *= -1
		neg = true
	}

	for range make([]struct{}, n) {
		a.Add(a, b)
		a, b = b, a
	}

	if neg && (n%2 == 0) {
		a = a.Neg(a)
	}
	return a
}

// This function calculates the n-th fibonacci number using the matrix method.
func FibMatrix(n int64) *big.Int {
	a, b := big.NewInt(1), big.NewInt(1)
	c, rc, tc := big.NewInt(1), big.NewInt(0), big.NewInt(0)
	d, rd := big.NewInt(0), big.NewInt(1)
	var neg bool

	if n < 0 {
		n = -n
		if n%2 == 0 {
			neg = true
		}
	}

	for n != 0 {
		if n&1 == 1 {
			tc = rc
			rc = new(big.Int).Add(new(big.Int).Mul(rc, a), new(big.Int).Mul(rd, c))
			rd = new(big.Int).Add(new(big.Int).Mul(tc, b), new(big.Int).Mul(rd, d))
		}

		ta := a
		tb := b
		tc = c
		a = new(big.Int).Add(new(big.Int).Mul(a, a), new(big.Int).Mul(b, c))
		b = new(big.Int).Add(new(big.Int).Mul(ta, b), new(big.Int).Mul(b, d))
		c = new(big.Int).Add(new(big.Int).Mul(c, ta), new(big.Int).Mul(d, c))
		d = new(big.Int).Add(new(big.Int).Mul(tc, tb), new(big.Int).Mul(d, d))

		n >>= 1
	}

	if neg {
		rc = rc.Neg(rc)
	}
	return rc
}
