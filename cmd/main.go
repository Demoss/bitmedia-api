package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello world")
}

func run() error {
	return nil
}
