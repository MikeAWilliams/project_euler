package util

func SumIntSlice(input []int) int {
	result := 0
	for _, item := range input {
		result += item
	}
	return result
}
