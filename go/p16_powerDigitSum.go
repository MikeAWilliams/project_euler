package main

import (
	"fmt"
	"math/big"
	"maw/util"
)

func main() {
	bigNumber := big.NewInt(2)
	bigNumber.Exp(big.NewInt(2), big.NewInt(1000), nil)

	result := util.SumDigitsBigInt(bigNumber)
	fmt.Println(result)
}
