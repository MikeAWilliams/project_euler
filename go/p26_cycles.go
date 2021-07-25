package main

import (
	"fmt"
	"math/big"
)

func getQuotients(max int, digits int) []string {
	quotients := make([]string, 1000)
	one := big.NewFloat(1.0)
	for d := 0; d < 1000; d++ {
		bigD := big.NewFloat(float64(d))
		quotient := big.NewFloat(1.0).Quo(one, bigD)
		textQuotient := quotient.Text('f', 40)
		quotients[d] = textQuotient
	}
	return quotients
}
func main() {
	quotients := getQuotients(1000, 40)
	fmt.Println(quotients)
	//0.1428571428571428492126926812488818541169
}
