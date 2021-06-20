package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	maxf := math.Sqrt(float64(num))
	max := int(maxf)
	if num%2 == 0 {
		return false
	}
	for i := 3; i <= max; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func sumPrimesBelow(max int) int {
	result := 2
	for candidate := 3; candidate < max; candidate += 2 {
		if isPrime(candidate) {
			result += candidate
		}
	}
	return result
}

func main() {
	fmt.Println(sumPrimesBelow(10))
	fmt.Println(sumPrimesBelow(2000000))
}
