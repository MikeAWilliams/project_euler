package main

import (
	"fmt"
	"maw/util"
)

func fibAsString(n int) string {
	fib := util.FibBig(n)
	fib.Text(10)
}

func digitsInFibN(n int) int {
	return len(fibAsString(n))
}

func main() {

	fmt.Println(digitsInFibN(200))
}
