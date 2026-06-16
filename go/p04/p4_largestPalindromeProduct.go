package main

import (
	"fmt"
	"strconv"
)

func reverse(source string) string {
	runes := []rune(source)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalendrome(num int) bool {
	forward := strconv.Itoa(num)
	backward := reverse(forward)
	return forward == backward
}

func bruteForceSearchForPalindromeProduct(max int) int {
	result := 0
	for i := max; i > 0; i-- {
		for j := max; j > 0; j-- {
			product := i * j
			if isPalendrome(product) {
				if product > result {
					result = product
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(bruteForceSearchForPalindromeProduct(100))
	fmt.Println(bruteForceSearchForPalindromeProduct(1000))
}
