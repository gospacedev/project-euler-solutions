package main

// Factorial returns the product of all positive integers less than or equal to a given positive integer
func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func p15() float64 {
	var (
		k float64 = 20
		n float64  = k * 2
	)

	// Solve for the amount all the of possible routes from (0, 0) to (20, 20) using combination formula
	ans := factorial(n) / (factorial(n-k) * factorial(k))
	return ans
}
