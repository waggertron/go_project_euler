package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumEvenFib(4000000)
}

func sumEvenFib(end int) int {
    x, y, sum := 1, 1, 0
    for sum < end {
        sum += x + y
        x, y = x+2*y, 2*x+3*y
    }
    return sum
}
