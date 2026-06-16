package main

import (
	"fmt"
	"maw/util"
)

func getPrimesBelowN(n int) []uint {
	filters := util.GetPrimesBelowFlags(n)
	return util.GetPrimes(filters)
}

func quadraticValue(a, b, n int) int {
	return n*n + a*n + b
}

func countPrimesForQuadratic(a, b int, primes []uint) int {
	n := 0
	val := quadraticValue(a, b, n)
	if val <= 1 {
		return 0
	}
	for util.IsPrime(uint(val), 0, primes) {
		n++
		val = quadraticValue(a, b, n)
		if val <= 1 {
			break
		}
	}
	return n
}

func findBestPrimeQuadratic(primes []uint) (int, int, int) {
	bestA := -1000
	bestB := -1000
	bestCount := -1
	for a := -1000; a < 1000; a++ {
		for b := -1000; b < 1000; b++ {
			count := countPrimesForQuadratic(a, b, primes)

			if count > bestCount {
				bestA = a
				bestB = b
				bestCount = count
			}
		}
	}
	return bestA, bestB, bestCount
}

func main() {
	const maxPrime = 10000
	primes := getPrimesBelowN(maxPrime)

	counta1b41 := countPrimesForQuadratic(1, 41, primes)
	fmt.Printf("a=1, b=41, count=%v\n", counta1b41)

	countaNeg79b1601 := countPrimesForQuadratic(-79, 1601, primes)
	fmt.Printf("a=-79, b=1601, count=%v\n", countaNeg79b1601)

	bestA, bestB, bestCount := findBestPrimeQuadratic(primes)
	fmt.Printf("bestA %v, bestB %v, bestCount %v, product %v\n", bestA, bestB, bestCount, bestA*bestB)
}
