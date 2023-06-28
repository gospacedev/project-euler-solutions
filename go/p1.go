package main

func p1() int {
	var n int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			n += i
		}
	}

	return n
}
