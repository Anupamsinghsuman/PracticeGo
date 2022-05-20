package main

import "fmt"

func main() {
	a := (42 == 45)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 46)

	fmt.Println(a, b, c, d)
	fmt.Println(a, b)
}
