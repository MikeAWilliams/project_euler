package main

import (
	"fmt"
	"maw/util"
	"strconv"
)

func fibAsString(n int) string {
	return strconv.Itoa(util.Fib(n))
}

func digitsInFibN(n int) int {
	return len(fibAsString(n))
}

func main() {

	fmt.Println(digitsInFibN(100))
}
