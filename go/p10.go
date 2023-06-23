package main

func p10() int {
	ans := 0
	for i := 0; ; i++ {
		if isPrime(i) { // isPrime is defined in p7.go
			ans += i
		} else if i >= 2000000 {
			break
		}
	}
	return ans
}
