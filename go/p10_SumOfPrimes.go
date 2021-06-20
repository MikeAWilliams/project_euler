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
