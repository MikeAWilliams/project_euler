package main

import (
	"fmt"
)

func collatz(n int) []int {
	result := []int{n}
	for currentNumber := n; currentNumber != 1; {
		if currentNumber%2 == 0 {
			currentNumber = currentNumber / 2
		} else {
			currentNumber = 3*currentNumber + 1
		}
		result = append(result, currentNumber)
	}

	return result
}

func lenCollatz(n int) int {
	return len(collatz(n))
}

func main() {
	maxLen := -1
	maxValue := -1
	for n := 2; n < 1000000; n++ {
		thisLen := lenCollatz(n)
		if thisLen > maxLen {
			maxValue = n
			maxLen = thisLen
		}
	}
	fmt.Printf("value %v, leng %v\n", maxValue, maxLen)
}
