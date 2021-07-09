package util

import (
	"fmt"
	"math"
	"sort"
)

func GetFactors(n int) []int {
	result := []int{1}
	if n > 1 {
		result = append(result, n)
	}
	stop := int(math.Sqrt(float64(n)))
	for i := 2; i <= stop; i++ {
		if n%i == 0 {
			result = append(result, i)
			result = append(result, n/i)
		}
	}
	sort.Ints(result)
	return result
}

func GetPrimeFactors(n int, allPrimesBelowN []int) []int {
	result := []int{}
	fmt.Printf("primes %v\n", allPrimesBelowN)
	for _, prime := range allPrimesBelowN {
		fmt.Printf("prime %v\n", prime)
		remainder := n % prime
		for remainder == 0 {
			result = append(result, prime)
			n = n / prime
			remainder = n % prime
		}
	}
	return result
}

func GetFactorsViaPrimes(n int, primeFactors []int) []int {
	return []int{}
}
