package main

import "fmt"

func main() {
	for c := 3; c <= 997; c++ {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				if isValid(a, b, c) && a+b+c == 1000 {
					fmt.Printf("a: %v, b: %v, c: %v\n", a, b, c)
					fmt.Println(a * b * c)
				}
			}
		}
	}
}

func isValid(a, b, c int) bool {
	return a < b && b < c && a*a+b*b == c*c
}
