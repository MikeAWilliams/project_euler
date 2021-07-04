package main

import (
	"bufio"
	"fmt"
	"log"
	"maw/util"
	"os"
)

func main() {
	file, err := os.Open("p13_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		thisLineData := []int{}
		for charIndex := 0; charIndex < len(line); charIndex++ {
			thisLineData = append(thisLineData, int(line[charIndex])-48)
		}
		data = append(data, thisLineData)
	}
	result := util.SumLargeInt(data)
	fmt.Println(result)
}
