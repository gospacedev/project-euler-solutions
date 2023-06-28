package main

func collatzSequenceLength(num int) int {
	var countTerms int = 0

	countTerms += 1

	for num > 1 {
		if num%2 == 0 {
			num = num / 2
			countTerms++
		} else {
			num = 3*num + 1
			countTerms++
		}
	}

	return countTerms
}

func p14() int {
	var (
		currentLength int
		longestLength int
	)

	for i := 0; i < 1000000; i++ {
		currentLength = collatzSequenceLength(i)

		if currentLength > longestLength {
			longestLength = currentLength
		}
		if longestLength == 525 {
			longestLength = i
		}
	}
	return longestLength
}
