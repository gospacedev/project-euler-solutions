package main

func findPythagoreanTriplet(sum int) (int, int, int) {
	for a := 1; a < sum/3; a++ {
		for b := a + 1; b < sum/2; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return 0, 0, 0
}

func p9() int {
	sum := 1000
	a, b, c := findPythagoreanTriplet(sum)
    product := a * b * c
    return product
}
