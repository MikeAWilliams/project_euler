package util

import (
	"math"
	"sort"
)

type FactorList struct {
	N          int
	FactorsOfN []int
}

type FactorListU struct {
	N          uint
	FactorsOfN []uint
}

func GetFactorsForAllNumbersBelowNUsingDivision(n int) []FactorList {
	result := make([]FactorList, n)
	for i := 1; i <= n; i++ {
		go func(index int) {
			result[index].N = index + 1
			result[index].FactorsOfN = GetFactors(index + 1)
		}(i - 1)
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

func isPrime(target uint, startIndex uint, data []uint) bool {
	stopIndex := uint(len(data) - 1)
	for startIndex <= stopIndex {
		index := (stopIndex + startIndex) / 2
		if data[index] == target {
			return true
		}
		if data[index] > target {
			stopIndex = index - 1
		} else {
			startIndex = index + 1
		}

	}
	return false
}

func GetPrimeFactors(n uint, primesToN []uint) []uint {
	result := []uint{}
	for primeIndex, prime := range primesToN {
		remainder := n % prime
		changedN := false
		for remainder == 0 {
			result = append(result, prime)
			n = n / prime
			remainder = n % prime
			changedN = true
		}
		if prime > n {
			break
		}
		if changedN && isPrime(n, uint(primeIndex), primesToN) {
			result = append(result, n)
			break
		}
	}
	return result
}

type primeFactorPower struct {
	prime  uint
	powers []uint
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

func (p *primeFactorPower) value() uint {
	return p.powers[p.index]
}

func computeFactor(primePowers []primeFactorPower) uint {
	result := uint(1)
	for _, prime := range primePowers {
		result *= prime.value()
	}
	return result
}

func countFactors(primePowers []primeFactorPower) int {
	result := 1
	for _, prime := range primePowers {
		result *= len(prime.powers)
	}
	return result
}

func buildPrimeFactorPower(prime uint, count int) primeFactorPower {
	result := primeFactorPower{prime: prime}
	power := uint(1)
	for i := 0; i <= count; i++ {
		result.powers = append(result.powers, power)
		power *= result.prime
	}
	return result
}

func getPrimeFactorPowers(numbers []uint) []primeFactorPower {
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

func GetFactorsViaPrimes(n uint, primeFactors []uint) []uint {
	if 1 == n {
		return []uint{1}
	}
	// https://math.stackexchange.com/questions/2782625/how-to-get-all-the-factors-of-a-number-using-its-prime-factorization
	countedPrimes := getPrimeFactorPowers(primeFactors)
	result := make([]uint, countFactors(countedPrimes))
	resultIndex := 0

	changingIndex := len(countedPrimes) - 1
	topChangingIndex := len(countedPrimes) - 1
	stop := false
	for !stop {
		result[resultIndex] = computeFactor(countedPrimes)
		resultIndex++
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
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func GetFactorsForAllNumbersBelowNUsingPrimes(n uint) []FactorListU {
	primeFlags := GetPrimesBelowFlags(int(n))
	primes := GetPrimes(primeFlags)
	result := make([]FactorListU, n)
	for i := uint(1); i <= n; i++ {
		go func(index uint) {
			result[index].N = index + 1
			result[index].FactorsOfN = GetFactorsViaPrimes(index+1, GetPrimeFactors(index+1, primes))
		}(i - 1)
	}
	return result
}
