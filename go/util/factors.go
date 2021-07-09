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

func GetPrimeFactors(n int, primesToN []int) []int {
	result := []int{}
	stop := n
	for _, prime := range primesToN {
		remainder := n % prime
		for remainder == 0 {
			result = append(result, prime)
			n = n / prime
			remainder = n % prime
		}
		if prime >= stop {
			break
		}
	}
	return result
}

func CountConsecutiveNumbers(numbers []int) map[int]int {
	result := map[int]int{}
	if 0 == len(numbers) {
		return result
	}
	count := 0
	lastNumber := numbers[0]
	for _, value := range numbers {
		if lastNumber == value {
			count++
		} else {
			result[lastNumber] = count
			lastNumber = value
			count = 1
		}
	}
	result[lastNumber] = count
	return result
}

func GetFactorsViaPrimes(n int, primeFactors []int) []int {
	result := []int{1}
	// https://math.stackexchange.com/questions/2782625/how-to-get-all-the-factors-of-a-number-using-its-prime-factorization
	fmt.Printf("Primes %v\n", primeFactors)
	countedPrimes := CountConsecutiveNumbers(primeFactors)
	for prime, count := range countedPrimes {
		toAdd := prime
		for i := 0; i < count; i++ {
			result = append(result, toAdd)
			toAdd *= prime
		}
	}
	sort.Ints(result)
	return result
}
