package main

import (
	"fmt"
	"math"
)

func main() {
	primes := getPrimes(2000000)
	sum := 0
	for _, val := range primes {
		sum += val
	}
	fmt.Println(sum)
}

// 142913828922
func getPrimes(end int) []int {
	nums := make([]int, 0, 100000)
	nums = append(nums, 2)
	for i := 3; i < end; i += 2 {
		nums = append(nums, i)
	}
	for i := 0; math.Pow(float64(nums[i]), 2.0) < float64(end); i++ {
		p := nums[i]
		leftOver := make([]int, 0, 10000)
		for j := i + 1; j < len(nums); j++ {
			val := nums[j]
			if val%p != 0 {
				leftOver = append(leftOver, val)
			}
		}
		nums = append(nums[:i+1], leftOver...)
	}
	return nums
}
