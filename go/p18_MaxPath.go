package main

import "fmt"

func maxTraverseGready(input [][]int) int {
	result := input[0][0]
	index := 0
	for row := 1; row < len(input); row++ {
		down := input[row][index]
		downRight := 0
		if index < len(input[row])-1 {
			downRight = input[row][index+1]
		}
		if down > downRight {
			result += down
		} else {
			result += downRight
			index++
		}

	}
	return result
}

func pickBestMoveInLookAhead(input [][]int, depth int, currentIndex int, row int, accumulator int) (int, int) {
	if 1 == depth {
		down := input[row][currentIndex]
		downRight := 0
		if currentIndex < len(input[row])-1 {
			downRight = input[row][currentIndex+1]
		}
		if down > downRight {
			return currentIndex, accumulator + down
		}
		return currentIndex + 1, accumulator + downRight
	}
	_, downResult := pickBestMoveInLookAhead(input, depth-1, currentIndex, row+1, accumulator+input[row][currentIndex])
	rightResult := 0
	if currentIndex < len(input[row])-1 {
		_, rightResult = pickBestMoveInLookAhead(input, depth-1, currentIndex+1, row+1, accumulator+input[row][currentIndex+1])
	}
	if downResult > rightResult {
		return currentIndex, downResult
	}
	return currentIndex + 1, rightResult
}

func maxTraverseLookAhead(input [][]int, depth int) int {
	result := input[0][0]
	index := 0
	for row := 1; row < len(input); row++ {
		searchDepth := depth
		if row+searchDepth > len(input) {
			searchDepth = len(input) - row
		}
		index, _ = pickBestMoveInLookAhead(input, searchDepth, index, row, 0)
		result += input[row][index]
	}
	return result
}

func main() {
	simpleIn := [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3},
	}
	fmt.Println(maxTraverseLookAhead(simpleIn, 3))
	hardIn := [][]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 03, 34},
		{88, 2, 77, 73, 07, 63, 67},
		{99, 65, 04, 28, 06, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23},
	}
	fmt.Println(maxTraverseGready(hardIn))
	//14 is max search depth makes this brute fource
	fmt.Println(maxTraverseLookAhead(hardIn, 14))
}
