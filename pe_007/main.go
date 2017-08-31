package main

import (
	"fmt"
	"math"
)

func main() {
	isPrime := primeChecker()
	count := 2
	end := 10001
	var k int
	for i := 1; ; i++ {
		k = 6 * i
		if isPrime(k - 1) {
			count++
			if count == end {
				fmt.Println(k - 1)
				break
			}
		}
		if isPrime(k + 1) {
			count++
			if count == end {
				fmt.Println(k + 1)
				break
			}
		}
	}
}

func primeChecker() func(int) bool {
	primes := make([]int, 0, 10001)
	primes = append(primes, 2, 3)
	return func(n int) bool {
		for _, v := range primes {
			if math.Sqrt(float64(n)) < float64(v) {
				break
			}
			if n%v == 0 {
				return false
			}
		}
		primes = append(primes, n)
		return true
	}
}
