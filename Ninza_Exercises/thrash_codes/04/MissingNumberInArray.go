package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 5, 6}
	sum := 0
	for _, k := range arr {
		sum = sum + k
	}
	fmt.Printf("The missing number is %v", (len(arr)+1)*(len(arr)+2)/2-sum)
}
