package main

import "fmt"

func main() {
	i := 10
	for i <= 100 {
		x := i % 4
		fmt.Println(x)
		i++
	}
}
