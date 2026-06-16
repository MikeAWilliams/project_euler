package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
			if date.Weekday() == time.Sunday {
				result++
			}
		}
	}
	fmt.Println(result)
}
