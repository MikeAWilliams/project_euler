package main

import (
	"fmt"
	"maw/util"
)

func sumPrimesBelow(max int) int {
	flags := util.GetPrimesBelowFlags(max)
	result := 0
	for index, flag := range flags {
		if flag {
			result += index + 2
		}
	}
	return result
}

func main() {
	result := (sumPrimesBelow(2000000))
	delta := 142913828922 - result
	fmt.Println(delta)
}
