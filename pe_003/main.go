package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxFactor(600851475143))
}
func findMaxFactor(num int64) int64 {
	factors := make([]int64, 0)
	for i := int64(2); i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			for num%i == 0 {
				num = num / i
			}
		}
	}
	return factors[len(factors)-1]
}
