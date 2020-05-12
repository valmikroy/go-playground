package main

import (
	"fmt"
	"strings"
)

func main() {

	// char by char looping
	input := "here is my string"

	for _, c := range input {
		fmt.Println(string(c))
	}

	fmt.Println("")

	// word by word looping
	for _, w := range strings.Fields(input) {
		fmt.Println(w)
	}
}
