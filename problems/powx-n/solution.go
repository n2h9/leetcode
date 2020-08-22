package solution

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		// n = -1 * n
		n = ^n + 1
	}
	if n == 1 {
		return x
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	halfPow := pow(x, n>>1)
	if n&1 == 0 {
		// n is even
		return halfPow * halfPow
	}
	return halfPow * halfPow * x
}
