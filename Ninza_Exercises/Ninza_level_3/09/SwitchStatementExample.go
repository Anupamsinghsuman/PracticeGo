package main

import "fmt"

func main() {
	switch "skiing" {
	case "skiing":
		fmt.Println("This is skiing")
	case "football":
		fmt.Println("This is football")
	default:
		fmt.Println("This is default case ")
	}
}
