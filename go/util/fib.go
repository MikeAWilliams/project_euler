package util

import (
	"math/big"
)

func Fib(n int) int {
	if 2 >= n {
		return 1
	}
	nm1 := 1
	nm2 := 1
	for index := 2; index < n; index++ {
		tmp := nm2 + nm1
		nm1 = nm2
		nm2 = tmp
	}
	return nm2
}

func FibBig(n int) *big.Int {
	if 2 >= n {
		return big.NewInt(1)
	}
	nm1 := big.NewInt(1)
	nm2 := big.NewInt(1)
	for index := 2; index < n; index++ {
		tmp := big.NewInt(0)
		tmp.Add(nm1, nm2)
		*nm1 = *nm2
		*nm2 = *tmp
	}
	return nm2
}
