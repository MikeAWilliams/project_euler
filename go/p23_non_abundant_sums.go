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
	fmt.Println("hello")
	properFactors := getProperFactorsForAllNumbersBelow(28123)
	//properFactors := getProperFactorsForAllNumbersBelow(50)
	abundantCandidates := getSumOfProperFactors(properFactors)
	abundants := getAbundant(abundantCandidates)
	fmt.Println("abundant")
	fmt.Println(abundants)
}
