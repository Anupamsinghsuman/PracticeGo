package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(bar(i) + bar(i-1))
	}
	os.Exit(1)
}
func bar(n int) {
	x := 0
	y := 1
	z := 0
	for i := 1; i <= n; i++ {
		z = x + y
		x = y
		y = z
	}
	return y
}
