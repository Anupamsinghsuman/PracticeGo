package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("This is true case ")
	case false:
		fmt.Println("This is false case ")
	default:
		fmt.Println("This is default case ")
	}
}
