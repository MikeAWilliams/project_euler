package main

import (
	"fmt"
)

func isDivisibleBy(num int, den int) bool {
	return num%den == 0
}

func getSumOfMultiples(max int) int {
	result := 0
	for i := 0; i < max; i++ {
		if isDivisibleBy(i, 3) || isDivisibleBy(i, 5) {
			result += i
		}
	}
	return result
}

func main() {
	fmt.Println(getSumOfMultiples(1000))
}
