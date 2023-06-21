package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func p4() int {
	var maxPalindrome int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			palindrome := i * j
			if strings.Compare(fmt.Sprintf("%d", palindrome), Reverse(fmt.Sprintf("%d", palindrome))) == 0 {
				if palindrome > maxPalindrome {
					maxPalindrome = palindrome
				}
			}
		}
	}
	return maxPalindrome
}