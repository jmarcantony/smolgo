package main

import (
    "fmt"
    "math"
)


func squareNum(i int) {
    fmt.Println(math.Pow(float64(i), 2))
}
func main() {
    for i := 0; i < 11; i++ {
        squareNum(i)
    }
}
