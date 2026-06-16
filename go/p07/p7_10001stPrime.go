package main

import (
	"fmt"
	"maw/util"
)

func findNthPrime(n int) int {
	flags := util.GetPrimesBelowFlags(n * 100)
	count := 0
	for index, flag := range flags {
		if flag {
			count++
			if count == n {
				return index
			}
		}
	}
	return -1
}

func main() {
	result := findNthPrime(10001)
	fmt.Printf("result %v\n", result)
}
