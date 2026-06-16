package main

import (
	"fmt"
	"maw/util"
)

func fibAsString(n int) string {
	fib := util.FibBig(n)
	return fib.Text(10)
}

func digitsInFibN(n int) int {
	return len(fibAsString(n))
}

func main() {

	min := 1000
	max := 8000
	fmt.Printf("%v, %v\n", digitsInFibN(min), digitsInFibN(max))
	searchResult := 0
	for min < max-1 {
		search := (max + min) / 2
		lenAtIndex := digitsInFibN(search)
		if 1000 == lenAtIndex {
			fmt.Printf("binary search found %v\n", search)
			searchResult = search
			break
		}
		if lenAtIndex > 1000 {
			max = search
		} else {
			min = search
		}
	}
	for i := searchResult; i > 0; i-- {
		digits := digitsInFibN(i)
		fmt.Printf("%v, %v\n", i, digits)
		if digits < 1000 {
			break
		}
	}
}
