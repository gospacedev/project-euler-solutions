package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	primes := []int{}
	isPrimeCounter := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			isPrimeCounter++
		} else if isPrimeCounter >= 10001 {
			break
		}
	}

	fmt.Println(primes)
}
