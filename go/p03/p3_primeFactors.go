package main

import (
	"fmt"
)

func primeFactors(number int) []int {
	result := []int{}
	for divisor := 2; number > 1; divisor++ {
		for number%divisor == 0 {
			result = append(result, divisor)
			number /= divisor
		}
	}

	return result
}

func main() {
	fmt.Println(primeFactors(2))
	fmt.Println(primeFactors(3))
	fmt.Println(primeFactors(4))
	fmt.Println(primeFactors(5))
	fmt.Println(primeFactors(9))
	fmt.Println(primeFactors(13195))
	fmt.Println(primeFactors(600851475143))
}
