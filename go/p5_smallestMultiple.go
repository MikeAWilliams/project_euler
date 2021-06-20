package main

import (
	"fmt"
)

func isDivisibleBy(num int, den int) bool {
	return num%den == 0
}

func isDivisibleByAllBelow(num int, maxDivisor int) bool {
	pass := true
	for divisor := 1; divisor <= maxDivisor; divisor++ {
		if !isDivisibleBy(num, divisor) {
			pass = false
			break
		}
	}
	return pass
}

func bruteForceSearchSmallestDivisibleByAllBelow(maxDivisor int) int {
	i := 1
	for ; i < 100000000000; i++ {
		if isDivisibleByAllBelow(i, maxDivisor) {
			return i
		}
	}
	return i
}

func main() {
	fmt.Println(isDivisibleByAllBelow(2520, 10))
	fmt.Println(bruteForceSearchSmallestDivisibleByAllBelow(20))
}
