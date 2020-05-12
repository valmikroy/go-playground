package main

import (
	"fmt"
	"strings"
)

func reverse(i string) {
	var c string
	if len(i) != 1 {
		c = string(i[0])
		reverse(string(i[1:]))
	} else {
		fmt.Print(string(i))
	}
	fmt.Print(c)

}

func reverseByWords(i []string) {

	var w string

	if len(i) != 1 {
		w = " " + strings.TrimSpace(i[0])
		reverseByWords(i[1:])
	} else {
		fmt.Print(strings.TrimSpace(i[0]))
	}
	fmt.Print(w)

}

func main() {

	input := "This is my string"

	//reverse(input)

	reverseByWords(strings.Fields(input))
}
