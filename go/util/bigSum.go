package util

func getMaxDigitsNumber(data [][]int) int {
	result := 0
	for _, slice := range data {
		thisLen := len(slice)
		if thisLen > result {
			result = thisLen
		}
	}
	return result
}

func SumLargeInt(data [][]int) []int {
	result := []int{}

	maxDigits := getMaxDigitsNumber(data)
	caryValue := 0
	for digitFromRight := 0; digitFromRight < maxDigits; digitFromRight++ {
		sumForThisDigit := caryValue
		caryValue = 0
		for _, slice := range data {
			thisDigitThisNumber := 0
			if len(slice) > digitFromRight {
				thisDigitThisNumber = slice[len(slice)-digitFromRight-1]
			}
			sumForThisDigit += thisDigitThisNumber
		}
		if sumForThisDigit > 9 {
			caryValue = sumForThisDigit / 10
			sumForThisDigit = sumForThisDigit % 10
		}
		result = append([]int{sumForThisDigit}, result...)
	}
	if caryValue > 0 {
		result = append([]int{caryValue}, result...)
	}

	return result
}
