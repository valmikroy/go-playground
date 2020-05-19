package main

import "fmt"

func appendVector() {

	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	c := append(a, b...)

	// combining vectors
	// [1 2 3 4 5 6 7 8]
	fmt.Println(c)

}

func main() {

	appendVector()
}
