package main

import (
    "fmt"
)

func printABCCandidates(){
    for a:=1; a < 1000; a++ {
        for b:=a + 1; b < 1000; b++ {
            c:= 1000 - b - a
            if c < b || c < a {
                break
            }
            if a * a + b*b == c * c {
                fmt.Printf("a,b,c %v %v %v\n", a, b, c)
                fmt.Printf("a2,b2,c2 %v %v %v\n", a*a, b*b, c*c)
                fmt.Printf("a2+b2 %v\n",a*a + b*b)
                fmt.Printf("a*b*c %v\n", a*b*c)
            }
        }
    }
}

func main(){
    printABCCandidates()
}