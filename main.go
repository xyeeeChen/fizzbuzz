package main

import (
	"fmt"
)

func main() {
	words := fizzBuzz(100)
	for _, word := range words {
		fmt.Println(word)
	}
}
