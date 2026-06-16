package main

import "fmt"

type kvp struct {
    value int
    representation string
}

func newKvp(v int, r string)kvp {
    return kvp{value:v, representation: r}
}

func numToWord(n int, values []kvp) string {
    result := ""
    for _, element := range values {
        if n >= element.value {
            result += element.representation
            n -= element.value
            if element.value < 1000 && element.value >= 100 && n > 0 {
                result += " and "
            }
        }
    }
    return result
}

func countLettersLikeTheyWant(value string)int {
    result :=0
    for i:=0; i<len(value); i++ {
        if ' ' != value[i] {
            result++
        }
    }
    fmt.Printf("%v, %v\n",value, result)
    return result
}

func main() {
    values := []kvp {
        newKvp(1000, "one thousand"),
        newKvp(900, "nine hundred"),
        newKvp(800, "eight hundred"),
        newKvp(700, "seven hundred"),
        newKvp(600, "six hundred"),
        newKvp(500, "five hundred"),
        newKvp(400, "four hundred"),
        newKvp(300, "three hundred"),
        newKvp(200, "two hundred"),
        newKvp(100, "one hundred"),
        newKvp(90, "ninety"),
        newKvp(80, "eighty"),
        newKvp(70, "seventy"),
        newKvp(60, "sixty"),
        newKvp(50, "fifty"),
        newKvp(40, "forty"),
        newKvp(30, "thirty"),
        newKvp(20, "twenty"),
        newKvp(19, "nineteen"),
        newKvp(18, "eighteen"),
        newKvp(17, "seventeen"),
        newKvp(16, "sixteen"),
        newKvp(15, "fifteen"),
        newKvp(14, "fourteen"),
        newKvp(13, "thirteen"),
        newKvp(12, "twelve"),
        newKvp(11, "eleven"),
        newKvp(10, "ten"),
        newKvp(9, "nine"),
        newKvp(8, "eight"),
        newKvp(7, "seven"),
        newKvp(6, "six"),
        newKvp(5, "five"),
        newKvp(4, "four"),
        newKvp(3, "three"),
        newKvp(2, "two"),
        newKvp(1, "one"), 
    }
    //fmt.Println(numToWord(1000, values))
    //fmt.Println(numToWord(999, values))
    //fmt.Println(numToWord(501, values))
    //fmt.Println(numToWord(516, values))
    //fmt.Println(countLettersLikeTheyWant(numToWord(342, values)))
    //fmt.Println(countLettersLikeTheyWant(numToWord(115, values)))
    result := 0
    for num:=1;num<=1000;num++ {
        result+=countLettersLikeTheyWant(numToWord(num, values))
    }
    fmt.Printf("\n\n%v\n", result)
}