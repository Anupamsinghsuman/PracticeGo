package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is Even Number\n", i)
		} else {
			fmt.Printf("%d is Odd Number\n", i)
		}
	}
}
