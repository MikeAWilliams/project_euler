package main

import (
	"fmt"
)

func isEven(n int) bool {
	return n%2 == 0
}

func sumEvenFib(max int) int {
	last := 1
	current := 1
	result := 0
	for current < max {
		fmt.Println(current)
		if isEven(current) {
			result += current
		}
		tmp := last
		last = current
		current = tmp + last
	}
	return result
}

func main() {
	fmt.Printf("\nresult is %v", sumEvenFib(4000000))
}
