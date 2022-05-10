package main

import "fmt"

func main() {
	x := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[0:4], x[5:]...)
	fmt.Println(x)
}
