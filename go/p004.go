//Unsolved
package main

import "fmt"

func isPal(n int) bool {
	r := 0
    for {
        if n > 0 {
            r = r*10 + n%10
            n = n / 10
        } else {
            break
        }
    }

	if n == r {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			ans := j * i

			if isPal(ans) {
				fmt.Println(ans)
			}
		}
	}
}