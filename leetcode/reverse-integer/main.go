/*
7. Reverse Integer
https://leetcode.com/problems/reverse-integer/
Given a 32-bit signed integer, reverse digits of an integer.

Input: 123
Output: 321

Input: -123
Output: -321

Input: 120
Output: 21

*/
package main

import (
	"fmt"
)

var base int = 10

func main() {
	i := 12345
	a, _ := div(i)
	fmt.Println(a)
}

func div(n int) (i int, j int) {
	quotient := n / base
	remainder := n % base

	if quotient > base {
		num, multiplier := div(quotient)
		glob := num + remainder*multiplier
		return glob, multiplier * base
	} else {
		glob := remainder*base + quotient
		return glob, base * base
	}
}

// 12345/10  = 1234 , 5
// 1234/10 = 123, 4
// 123/10 = 12, 3
// 12/10 = 1, 2

// 1 * 1 + 2 * 10
// 3 * 100
// 4 * 1000
// 5 * 10000

// divide number by base , take out quotient and remainder
// if quotient is greater than recurse
// when quoitent is less than base then multiply quoitent by 1 and remainder by base

// return a number and next multiplier
