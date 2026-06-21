// fun fact this takes 0.29 s in go
// in python it takes 7.97
package main

import "fmt"

func getNext(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n = n / 10
	}
	return sum
}

func ends89(n int) bool {
	for n != 1 && n != 89 {
		n = getNext(n)
	}
	return n == 89
}

func main() {
	count := 0
	for i := 1; i < 10000000; i++ {
		if ends89(i) {
			count += 1
		}
	}
	fmt.Println(count)
}
