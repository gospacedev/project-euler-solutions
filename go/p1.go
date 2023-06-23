package main

func sum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func p1() int {
	var a []int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			a = append(a, i)
		}
	}

	ans := sum(a)

	return ans
}
