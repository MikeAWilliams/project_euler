package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	maxf := math.Sqrt(float64(num))
	max := int(maxf)
	for i := 2; i <= max; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findNthPrime(n int) int {
	candidate := 2
	for found := 0; found < n; {
		if isPrime(candidate) {
			found += 1
		}
		candidate++
	}
	return candidate - 1
}

func main() {
	fmt.Println(findNthPrime(6))
	fmt.Println(findNthPrime(10001))
}