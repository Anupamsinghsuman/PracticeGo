package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

// when we use log.Fatalln deferred function do not run
