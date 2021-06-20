package main

import (
	"fmt"
)

func isPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
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
