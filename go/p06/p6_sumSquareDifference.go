package main

import (
	"fmt"
)

func sumSquare(maxNum int) int {
	result := 0
	for i := 1; i <= maxNum; i++ {
		result += i * i
	}
	return result
}

func squareSum(maxNum int) int {
	result := 0
	for i := 1; i <= maxNum; i++ {
		result += i
	}
	return result * result
}

func main() {
	fmt.Println(sumSquare(10))
	fmt.Println(squareSum(10))

	squareSum100 := squareSum(100)
	sumSquare100 := sumSquare(100)
	fmt.Println(squareSum100 - sumSquare100)
}
