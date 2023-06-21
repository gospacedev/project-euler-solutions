package main

func primeFactors(n int) []int {

	var v []int

	for n%2 == 0 {
		v = append(v, n)
		n = n / 2
	}

	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			v = append(v, n)
			n = n / i
		}
	}

	if n > 2 {
		v = append(v, n)
	}

	return v
}

func p3() []int {
	ans := primeFactors(600851475143)

	return ans
}
