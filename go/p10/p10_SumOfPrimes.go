package main

import (
	"fmt"
	"maw/util"
)

func sumPrimesBelow(max int) uint {
	flags := util.GetPrimesBelowFlags(max)
	primes := util.GetPrimes(flags)
	result := uint(0)
	for _, prime := range primes {
		result += prime
	}
	return result
}

func main() {
	result := (sumPrimesBelow(2000000))
	delta := 142913828922 - result
	fmt.Println(delta)
}
