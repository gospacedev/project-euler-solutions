package main

func p2() int {
	ans := 0
	x := 1
	y := 2
	for x < 4000000 {
		if x%2 == 0 {
			ans += x
		}

		x, y = y, x+y
	}

	return ans
}
