package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	fileContentsRaw, err := ioutil.ReadFile("p022_names.txt")
	if nil != err {
		panic(err)
	}
	fileString := string(fileContentsRaw)
	fileString = strings.ReplaceAll(fileString, "\"", "")
	names := strings.Split(fileString, ",")
	sort.Strings(names)
	fmt.Println(names)
}
