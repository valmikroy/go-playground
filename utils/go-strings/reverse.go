package main

import (
	"fmt"
	"strings"
)

func reverseByChars(i string) {
	var c string
	if len(i) != 1 {
		c = string(i[0])
		reverseByChars(string(i[1:]))
	} else {
		fmt.Print(string(i))
	}
	fmt.Print(c)

}

func reverseByCharWithSlice(i string) {
	cnt := len(i)
	for j := cnt - 1; j >= 0; j-- {
		fmt.Print(string(i[j]))
	}
	fmt.Println()
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

	//reverseByChars(input)
	reverseByCharWithSlice(input)

	fmt.Println()

	//reverseByWords(strings.Fields(input))
}
