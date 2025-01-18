package main

func collatzSequenceLength(num int) int {
	countTerms := 1

	for num > 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}
		countTerms++
	}

	return countTerms
}

func p14() int {
	var (
		currentLength int
		longestLength int
		startNum int
	)

	for i := 0; i < 1000000; i++ {
		currentLength = collatzSequenceLength(i)

		if currentLength > longestLength {
			longestLength = currentLength
			startNum = i
		}

	}
	return startNum
}
