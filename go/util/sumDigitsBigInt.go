package util

import "math/big"

func SumDigitsBigInt(bigNumber *big.Int) int64 {
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
	return result
}
