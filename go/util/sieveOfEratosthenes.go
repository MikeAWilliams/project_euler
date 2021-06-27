package util

func getInitialSlice(n int) []bool {
	result := make([]bool, n-1)
	for i := range result {
		result[i] = true
	}
	return result
}

func GetPrimesBelow(n int) []bool {
	candidates := getInitialSlice(n)
	candidateIndex := 0
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

func CountTrue(in []bool) int {
	result := 0
	for _, item := range in {
		if item {
			result++
		}
	}
	return result
}
