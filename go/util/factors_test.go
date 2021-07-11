package util_test

import (
	"fmt"
	"maw/util"
	"testing"
)

func requireEqualSlices(t *testing.T, expected, recieved []int) {
	if len(expected) != len(recieved) {
		fmt.Printf("expected %v\n", expected)
		fmt.Printf("recieved %v\n", recieved)
		t.Fail()
		return
	}
	for index, expectedValue := range expected {
		if expectedValue != recieved[index] {
			fmt.Printf("expected %v\n", expected)
			fmt.Printf("recieved %v\n", recieved)
			t.Fail()
			return
		}
	}
}

func requireEqualMaps(t *testing.T, expected, recieved map[int]int) {
	if len(expected) != len(recieved) {
		fmt.Printf("expected %v\n", expected)
		fmt.Printf("recieved %v\n", recieved)
		t.Fail()
		return
	}
	for eKey, eValue := range expected {
		rValue, ok := recieved[eKey]
		if !ok || rValue != eValue {
			fmt.Printf("expected %v\n", expected)
			fmt.Printf("recieved %v\n", recieved)
			t.Fail()
			return
		}
	}
}

func TestFactors(t *testing.T) {
	requireEqualSlices(t, []int{1}, util.GetFactors(1))
	requireEqualSlices(t, []int{1, 3}, util.GetFactors(3))
	requireEqualSlices(t, []int{1, 2, 3, 6}, util.GetFactors(6))
	requireEqualSlices(t, []int{1, 2, 4, 7, 14, 28}, util.GetFactors(28))
	requireEqualSlices(t, []int{1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500}, util.GetFactors(500))
}

func BenchmarkFactorsByDivision(b *testing.B) {
	var result []util.FactorList
	for n := 0; n < b.N; n++ {
		result = util.GetFactorsForAllNumbersBelowNUsingDivision(28123)
	}
	result[0].N = 1
}

func TestGetPrimeFactors(t *testing.T) {
	primeFlagsBelow1000 := util.GetPrimesBelowFlags(1000)
	primesBelow1000 := util.GetPrimes(primeFlagsBelow1000)
	requireEqualSlices(t, []int{}, util.GetPrimeFactors(1, primesBelow1000))
	requireEqualSlices(t, []int{3}, util.GetPrimeFactors(3, primesBelow1000))
	requireEqualSlices(t, []int{5}, util.GetPrimeFactors(5, primesBelow1000))
	requireEqualSlices(t, []int{2, 2, 2, 3}, util.GetPrimeFactors(24, primesBelow1000))
	requireEqualSlices(t, []int{2, 2, 3, 3, 3, 3}, util.GetPrimeFactors(324, primesBelow1000))
	requireEqualSlices(t, []int{3, 3, 3, 37}, util.GetPrimeFactors(999, primesBelow1000))
}

func BenchmarkGetPrimeFactors(b *testing.B) {
	primeFlagsBelow1000 := util.GetPrimesBelowFlags(1000)
	primesBelow1000 := util.GetPrimes(primeFlagsBelow1000)
	var result []int
	for n := 0; n < b.N; n++ {
		result = util.GetPrimeFactors(999, primesBelow1000)
	}
	result[0] = 1
}

func BenchmarkGetFactors(b *testing.B) {
	var result []int
	for n := 0; n < b.N; n++ {
		result = util.GetFactors(999)
	}
	result[0] = 1
}

func TestGetFactorsViaPrimes(t *testing.T) {
	primeFlagsBelow1000 := util.GetPrimesBelowFlags(1000)
	primesBelow1000 := util.GetPrimes(primeFlagsBelow1000)
	requireEqualSlices(t, []int{1}, util.GetFactorsViaPrimes(1, util.GetPrimeFactors(1, primesBelow1000)))
	requireEqualSlices(t, []int{1, 3}, util.GetFactorsViaPrimes(3, util.GetPrimeFactors(3, primesBelow1000)))
	requireEqualSlices(t, []int{1, 2, 3, 6}, util.GetFactorsViaPrimes(6, util.GetPrimeFactors(6, primesBelow1000)))
	requireEqualSlices(t, []int{1, 2, 4, 7, 14, 28}, util.GetFactorsViaPrimes(28, util.GetPrimeFactors(28, primesBelow1000)))
	requireEqualSlices(t, []int{1, 2, 3, 4, 6, 9, 12, 18, 27, 36, 54, 81, 108, 162, 324}, util.GetFactorsViaPrimes(324, util.GetPrimeFactors(324, primesBelow1000)))
	requireEqualSlices(t, []int{1, 2, 3, 4, 5, 6, 7, 10, 12, 14, 15, 20, 21, 28, 30, 35, 42, 60, 70, 84, 105, 140, 210, 420}, util.GetFactorsViaPrimes(420, util.GetPrimeFactors(420, primesBelow1000)))
	requireEqualSlices(t, []int{1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500}, util.GetFactorsViaPrimes(500, util.GetPrimeFactors(500, primesBelow1000)))
}

func BenchmarkFactorsByPrimes(b *testing.B) {
	var result []util.FactorList
	for n := 0; n < b.N; n++ {
		result = util.GetFactorsForAllNumbersBelowNUsingPrimes(28123)
	}
	result[0].N = 1
}
