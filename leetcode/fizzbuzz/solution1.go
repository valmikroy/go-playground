/*
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and
for the multiples of five output “Buzz”. For numbers which are multiples of
both three and five output “FizzBuzz”.
*/
package main

import "fmt"

func main() {

	for i := 1; i <= 15; i++ {

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizBuzz")
		case i%3 == 0 && i%5 != 0:
			fmt.Println("Fiz")
		case i%3 != 0 && i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}

	}
}
