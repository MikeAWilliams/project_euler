package main

import (
	"fmt"
	"math/big"
)

func main() {
	min := 2
	//max := 5
	max := 100
	values := make(map[string]struct{})

	var p big.Int
	for a := min; a <= max; a++ {
		for b := min; b <= max; b++ {
			p.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			values[p.String()] = struct{}{}
		}
	}
	fmt.Println(len(values))
}
