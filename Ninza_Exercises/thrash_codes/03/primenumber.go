package main

import (
	"fmt"
	"math"
)

const (
	x = 45
	i = 2.00
)

func main() {
	res := math.Sqrt(x)
	for i <= res {
		if x%i == 0 {
			fmt.Println("not a prime number")
		} else {
			fmt.Println("Prime number")
		}
		i = i + 1.00
	}
}
