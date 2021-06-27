package util

func getInitialFlags(n int) []bool {
	result := make([]bool, n-1)
	result[0] = true
	for i := 1; i < len(result); i += 2 {
		result[i] = true
	}
	return result
}

func GetPrimesBelowFlags(n int) []bool {
	candidates := getInitialFlags(n)
	candidateIndex := 1
	candidate := candidateIndex + 2
	for candidate*candidate < n {
		for indexToRemove := candidateIndex + candidate; indexToRemove < len(candidates); indexToRemove += candidate {
			candidates[indexToRemove] = false
		}
		candidateIndex++
		candidate = candidateIndex + 2
	}
	return candidates
}

func CountPrimeFlags(flags []bool) int {
	result := 0
	for _, item := range flags {
		if item {
			result++
		}
	}
	return result
}

func GetPrimes(flags []bool) []int {
	result := []int{}
	for index, flag := range flags {
		if flag {
			result = append(result, index+2)
		}
	}
	return result
}
