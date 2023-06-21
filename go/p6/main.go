package main

import (
	"fmt"
	"math"
)

func main() {
	sumOfSqaures := 0
	squareOfSums := 0

	for i := 0; i <= 100; i++ {
		sumOfSqaures += int(math.Pow(float64(i), 2))

		squareOfSums += i
	}

	squareOfSums = int(math.Pow(float64(squareOfSums), 2))

	n := squareOfSums - sumOfSqaures

	fmt.Println(n)
}
