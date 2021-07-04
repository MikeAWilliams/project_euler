package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigNumber := big.NewInt(2)
	bigNumber.Exp(big.NewInt(2), big.NewInt(1000), nil)

	ten := big.NewInt(10)
	junk := big.NewInt(10)
	var result int64
	result = 0
	keepGointg := true
	for keepGointg {
		_, m := bigNumber.DivMod(bigNumber, ten, junk)
		result += m.Int64()
		digits := len(bigNumber.Text(10))
		if 1 == digits {
			keepGointg = (bigNumber.Int64() > 0)
		}
	}
	fmt.Println(result)
}
