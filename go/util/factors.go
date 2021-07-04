package util

import (
	"math"
	"sort"
)

func GetFactors(n int) []int {
	result := []int{1}
	if n > 1 {
		result = append(result, n)
	}
	stop := int(math.Sqrt(float64(n)))
	for i := 2; i <= stop; i++ {
		if n%i == 0 {
			result = append(result, i)
			result = append(result, n/i)
		}
	}
	sort.Ints(result)
	return result
}
