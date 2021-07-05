package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func wordScore(word string) int {
	result := 0
	for i := 0; i < len(word); i++ {
		result += int(word[i]-'A') + 1
	}
	return result
}

func main() {
	fileContentsRaw, err := ioutil.ReadFile("p022_names.txt")
	if nil != err {
		panic(err)
	}
	fileString := string(fileContentsRaw)
	fileString = strings.ReplaceAll(fileString, "\"", "")
	names := strings.Split(fileString, ",")
	sort.Strings(names)

	result := 0
	for index, name := range names {
		result += (index + 1) * wordScore(name)
	}
	fmt.Println(result)
}
