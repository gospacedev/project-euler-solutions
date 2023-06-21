package main

import (
	"math"
)

func p6() int {
	sumOfSqaures := 0
	squareOfSums := 0

	for i := 0; i <= 100; i++ {
		sumOfSqaures += int(math.Pow(float64(i), 2))

		squareOfSums += i
	}

	squareOfSums = int(math.Pow(float64(squareOfSums), 2))

	ans := squareOfSums - sumOfSqaures

	return ans
}
