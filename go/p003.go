package main

import (
	"fmt"
	//"math"
)

func primeFactors(n int) []int {

	var v []int
	
	// Print the number of 2s that divide n
	for n%2 == 0 {
		v = append(v, n)
		n = n / 2
	}

	// n must be odd at this point. So we can skip
	// one element (Note i = i + 2)
	for  i := 3; i*i <= n; i = i + 2 {
		// While i divides n, print i and divide n
		for n%i == 0 {
			v = append(v, n)
			n = n / i
		}
	}

	// This condition is to handle the case when n
	// is a prime number greater than 2
	if n > 2 {
		v = append(v, n)
	}

	return v
}

func main() {
	ans := primeFactors(600851475143)

	fmt.Println(ans)
}