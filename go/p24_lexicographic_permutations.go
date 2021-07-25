package main

import (
	"fmt"
	"sort"
	"strconv"
)

func heapRecursive(source []int, last int, result *[][]int) {
	if 1 == last {
		thisResult := make([]int, len(source))
		copy(thisResult, source)
		*result = append(*result, thisResult)
		return
	}
	for i := 0; i < last; i++ {
		heapRecursive(source, last-1, result)

		if 1 == last%2 {
			tmp := source[0]
			source[0] = source[last-1]
			source[last-1] = tmp
		} else {
			tmp := source[i]
			source[i] = source[last-1]
			source[last-1] = tmp
		}
	}
}
func heapsAlgorithm(source []int) [][]int {
	var result [][]int
	heapRecursive(source, len(source), &result)
	return result
}

func intSliceToString(input []int) string {
	var result string
	for _, digit := range input {
		result += strconv.Itoa(digit)
	}
	return result
}

func intSlicesToStringSlice(input [][]int) []string {
	result := make([]string, len(input))
	for index, slice := range input {
		result[index] = intSliceToString(slice)
	}
	return result
}

func main() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	permutations := heapsAlgorithm(input)
	stringPermulations := intSlicesToStringSlice(permutations)
	sort.Strings(stringPermulations)
	fmt.Println(stringPermulations[1000000])
	fmt.Println(stringPermulations[1000001])
}
