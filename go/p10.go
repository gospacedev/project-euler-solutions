package main

func p10() int {
	ans := 0
	for i := 0; i <= 2000000; i++ {
		if isPrime(i) { // isPrime is defined in p7.go
			ans += i
		}
	}
	return ans
}
