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
				return index + 2
			}
		}
	}
	return -1
}

func main() {
	result := findNthPrime(10001)
	fmt.Printf("result %v\n", result)
	fmt.Printf("delta %v", 104743-result)
}
