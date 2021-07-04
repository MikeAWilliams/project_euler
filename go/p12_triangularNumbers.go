package main

import (
	"fmt"
	"maw/util"
)

func getTraingleNumber(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func main() {
	n := 1
	for {
		triangle := getTraingleNumber(n)
		factors := util.GetFactors(triangle)
		if len(factors) > 500 {
			fmt.Printf("n %v t %v\n", n, triangle)
			break
		}
		n++
	}

}
