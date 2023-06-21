package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	smallestMultiple := 1
	for i := 2; i <= 20; i++ {
		smallestMultiple = lcm(smallestMultiple, i)
	}
	fmt.Println(smallestMultiple)
}

