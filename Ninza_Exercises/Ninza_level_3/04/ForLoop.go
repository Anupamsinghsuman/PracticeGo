package main

import "fmt"

func main() {
	i := 2000
	for {
		if i == 2023 {
			break
		}
		fmt.Println(i)
		i++
	}
}
