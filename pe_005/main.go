package main

import (
	"fmt"
)

func main() {
	factorMap := makeFactorMap()
	for i := 1; i <= 20; i++ {
		fmt.Println(factorMap(i))
	}
	res := factorMap(1)
	product := 1
	for _, factor := range res {
		product *= factor
	}
	fmt.Println(product)
}

func primeFactors(n int) []int {
	res := make([]int, 0)
	if n < 2 {
		return res
	}
	for i := 2; n >= i; i++ {
		for n%i == 0 {
			res = append(res, i)
			n = n / i
		}
	}
	return res
}
func makeFactorMap() func(int) []int {
	factors := make(map[int]int)
	return func(n int) []int {
		nFactors := primeFactors(n)
		nFactorCount := make(map[int]int)
		for _, val := range nFactors {
			nFactorCount[val]++
		}
		for k, v := range nFactorCount {
			if v > factors[k] {
				factors[k] = v
			}
		}
		allFactors := make([]int, 0)
		for k, v := range factors {
			for ; v > 0; v-- {
				allFactors = append(allFactors, k)
			}
		}

		return allFactors
	}
}
