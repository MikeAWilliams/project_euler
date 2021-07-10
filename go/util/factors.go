package util

import (
	"math"
	"sort"
)

type FactorList struct {
	N          int
	FactorsOfN []int
}

func GetFactorsForAllNumbersBelowNUsingDivision(n int) []FactorList {
	result := make([]FactorList, n)
	for i := 1; i <= n; i++ {
		result[i-1].N = i
		result[i-1].FactorsOfN = GetFactors(i)
	}
	return result
}

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

type primeFactorPower struct {
	prime  int
	powers []int
	index  int
}

func (p *primeFactorPower) increment() bool {
	if p.index < len(p.powers)-1 {
		p.index++
		return false
	}
	p.index = 0
	return true
}

func (p *primeFactorPower) value() int {
	return p.powers[p.index]
}

func computeFactor(primePowers []primeFactorPower) int {
	result := 1
	for _, prime := range primePowers {
		result *= prime.value()
	}
	return result
}

func buildPrimeFactorPower(prime int, count int) primeFactorPower {
	result := primeFactorPower{prime: prime}
	power := 1
	for i := 0; i <= count; i++ {
		result.powers = append(result.powers, power)
		power *= result.prime
	}
	return result
}

func getPrimeFactorPowers(numbers []int) []primeFactorPower {
	result := []primeFactorPower{}
	if 0 == len(numbers) {
		return result
	}
	count := 0
	lastValue := numbers[0]
	for _, value := range numbers {
		if lastValue == value {
			count++
		} else {
			result = append(result, buildPrimeFactorPower(lastValue, count))
			lastValue = value
			count = 1
		}
	}
	result = append(result, buildPrimeFactorPower(lastValue, count))
	return result
}

func GetFactorsViaPrimes(n int, primeFactors []int) []int {
	if 1 == n {
		return []int{1}
	}
	result := []int{}
	// https://math.stackexchange.com/questions/2782625/how-to-get-all-the-factors-of-a-number-using-its-prime-factorization
	countedPrimes := getPrimeFactorPowers(primeFactors)
	changingIndex := len(countedPrimes) - 1
	topChangingIndex := len(countedPrimes) - 1
	stop := false
	for !stop {
		result = append(result, computeFactor(countedPrimes))
		roll := countedPrimes[changingIndex].increment()
		for roll {
			if topChangingIndex == changingIndex {
				if 0 == topChangingIndex {
					stop = true
					break
				}
				topChangingIndex--
			}
			roll = countedPrimes[changingIndex-1].increment()
			if roll {
				changingIndex--
			} else {
				changingIndex = len(countedPrimes) - 1
			}
		}
	}
	sort.Ints(result)
	return result
}

func GetFactorsForAllNumbersBelowNUsingPrimes(n int) []FactorList {
	primeFlags := GetPrimesBelowFlags(n)
	primes := GetPrimes(primeFlags)
	result := make([]FactorList, n)
	for i := 1; i <= n; i++ {
		result[i-1].N = i
		result[i-1].FactorsOfN = GetFactorsViaPrimes(i, GetPrimeFactors(i, primes))
	}
	return result
}
