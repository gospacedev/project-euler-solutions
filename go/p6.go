package main

import (
	"math"
)

func p6() int {
	sumOfSquares := 0
	squareOfSums := 0

	for i := 0; i <= 100; i++ {
		sumOfSquares += int(math.Pow(float64(i), 2))

		squareOfSums += i
	}

	squareOfSums = int(math.Pow(float64(squareOfSums), 2))

	ans := squareOfSums - sumOfSquares

	return ans
}
