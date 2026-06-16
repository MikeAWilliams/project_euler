package main

import (
	"bufio"
	"fmt"
	"log"
	"maw/util"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("p13/data.txt")
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

	answer := ""
	for _, digit := range result[:10] {
		answer += strconv.Itoa(digit)
	}
	fmt.Println(answer)
}
