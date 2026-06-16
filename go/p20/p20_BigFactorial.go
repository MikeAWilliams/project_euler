package main

import (
	"fmt"
	"math/big"
	"maw/util"
)

func bigFactorial(n int64) *big.Int {
	result := big.NewInt(n)
	for n--; n > 0; n-- {
		result = result.Mul(result, big.NewInt(n))
	}
	return result
}

func sumDigitsBigFactorial(n int64) int64 {
	bigNumber := bigFactorial(n)
	return util.SumDigitsBigInt(bigNumber)
}

func main() {
	fmt.Println(sumDigitsBigFactorial(int64(10)))
	fmt.Println(sumDigitsBigFactorial(int64(100)))
}
