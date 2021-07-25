package main

import (
	"fmt"
	"maw/util"
	"sync"
)

func getProperFactorsForAllNumbersBelow(n int) []util.FactorList {
	factors := util.GetFactorsForAllNumbersBelowNUsingDivision(n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := range factors {
		go func(index int) {
			factors[index].FactorsOfN = factors[index].FactorsOfN[:len(factors[index].FactorsOfN)-1]
			wg.Done()
		}(i)
	}
	wg.Wait()
	return factors
}

type abundantCheck struct {
	n                  int
	sumOfProperFactors int
}

func sumSlice(in []int) int {
	result := 0
	for _, val := range in {
		result += val
	}
	return result
}

func getSumOfProperFactors(factors []util.FactorList) []abundantCheck {
	result := make([]abundantCheck, len(factors))
	var wg sync.WaitGroup
	wg.Add(len(factors))
	for i := range factors {
		go func(index int) {
			result[index].n = factors[index].N
			result[index].sumOfProperFactors = sumSlice(factors[index].FactorsOfN)
			wg.Done()
		}(i)
	}
	return result
}

func getAbundant(input []abundantCheck) []abundantCheck {
	var result []abundantCheck
	for _, candidate := range input {
		if candidate.n < candidate.sumOfProperFactors {
			result = append(result, candidate)
		}
	}
	return result
}

func main() {
	const maxSearch = 28123
	properFactors := getProperFactorsForAllNumbersBelow(maxSearch)
	//properFactors := getProperFactorsForAllNumbersBelow(50)
	abundantCandidates := getSumOfProperFactors(properFactors)
	abundants := getAbundant(abundantCandidates)
	isSumOfAbundant := make([]bool, maxSearch)
	for _, abundant1 := range abundants {
		for _, abundant2 := range abundants {
			sum := abundant1.n + abundant2.n
			if sum < maxSearch {
				isSumOfAbundant[sum] = true
			} else {
				break
			}
		}
	}
	sumCantBeAbundantSum := 0
	for i := 1; i < maxSearch; i++ {
		if !isSumOfAbundant[i] {
			sumCantBeAbundantSum += i
		}
	}
	fmt.Println(sumCantBeAbundantSum)
}
