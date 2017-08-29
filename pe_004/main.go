package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(largestPalinProduct(3))

}
func largestPalinProduct(digits int) int {
	start := math.Pow10(digits) - 1
	max := 0
	for i := start; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			product := int(i * j)
			if product > max && isNumPalin(product) {
				max = product
			}
		}
	}
	return max
}
func isNumPalin(n int) bool {
	nStr := strconv.Itoa(n)
	res := ""
	for _, r := range nStr {
		res = string(r) + res
	}
	return res == nStr

}
