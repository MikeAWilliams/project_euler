package main

import (
	"fmt"
	"maw/util"
)

func isAmiable(n int, properFactors map[int][]int) (bool, int) {
	divisorsN := properFactors[n]
	sum := util.SumIntSlice(divisorsN)
	if sum == n {
		return false, -1
	}
	divisorsSum := properFactors[sum]
	sumDivisorsSum := util.SumIntSlice(divisorsSum)
	return sumDivisorsSum == n, sum
}

func main() {
	properFactors := map[int][]int{}
	for n := 1; n < 10000; n++ {
		thisFactors := util.GetFactors(n)
		thisProperFactors := thisFactors[:len(thisFactors)-1]
		properFactors[n] = thisProperFactors
	}
	amiable := map[int]int{}
	for n := 1; n < 10000; n++ {
		other, present := amiable[n]
		if present {
			continue
		}
		nIsAmiable, other := isAmiable(n, properFactors)
		if nIsAmiable {
			amiable[n] = other
			amiable[other] = n
		}
	}

	result := 0
	for k, _ := range amiable {
		result += k
	}
	fmt.Println(result)
}
