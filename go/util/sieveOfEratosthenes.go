package util

func getInitialFlags(n int) []bool {
	result := make([]bool, n+1)
	result[2] = true
	for i := 3; i < len(result); i += 2 {
		result[i] = true
	}
	return result
}

func GetPrimesBelowFlags(n int) []bool {
	candidates := getInitialFlags(n)
	candidate := 3
	for candidate*candidate < n {
		for indexToRemove := candidate * candidate; indexToRemove < len(candidates); indexToRemove += candidate * 2 {
			candidates[indexToRemove] = false
		}
		for candidate += 2; !candidates[candidate]; candidate += 2 {
		}
	}
	return candidates
}

func CountPrimeFlags(flags []bool) int {
	result := 1
	for index := 3; index < len(flags); index += 2 {
		if flags[index] {
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
