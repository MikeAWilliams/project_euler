package main
import "fmt"

func main() {
    for d:=0.0; d < 1000.0; d++{
        result := 1.0/d
        fmt.Printf("%v, %v\n",d, result)
    }

}