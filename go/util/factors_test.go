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

func TestFactors(t *testing.T) {
	requireEqualSlices(t, []int{1}, util.GetFactors(1))
	requireEqualSlices(t, []int{1, 3}, util.GetFactors(3))
	requireEqualSlices(t, []int{1, 2, 3, 6}, util.GetFactors(6))
	requireEqualSlices(t, []int{1, 2, 4, 7, 14, 28}, util.GetFactors(28))
	requireEqualSlices(t, []int{1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500}, util.GetFactors(500))
}

func TestGetPrimeFactors(t *testing.T) {
	primeFlagsBelow1000 := util.GetPrimesBelowFlags(1000)
	primesBelow1000 := util.GetPrimes(primeFlagsBelow1000)
	requireEqualSlices(t, []int{2, 2, 2, 3}, util.GetPrimeFactors(24, primesBelow1000))
}
