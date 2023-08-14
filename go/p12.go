package main

import "math"

func triangularNumber(n int) int {
	return n * (n + 1) / 2
}

func countDivisors(n int) int {
	count := 0
	sqrt := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			count += 2 // Count both divisors: i and n/i
		}
	}

	// If n is a perfect square, only count one divisor
	if sqrt*sqrt == n {
		count--
	}

	return count
}

func p12() int {
	var ans int
	for i := 1; ; i++ {
		if countDivisors(triangularNumber(i)) >= 500 {
			ans += triangularNumber(i)
			return ans
		}
	}
}
