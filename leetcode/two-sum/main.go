/*
1. Two Sum
Source: https://leetcode.com/problems/two-sum/
Given an array of integers, return indices of the two numbers such that they
add up to a specific target.  You may assume that each input would have exactly
one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

package main

import (
	"fmt"
)

func main() {

	in := []int{2, 7, 11, 15}
	target := 9
	var sum int

Loop:
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in); j++ {
			sum = in[i] + in[j]
			if sum == target {
				fmt.Printf("%d and %d\n", i, j)
				break Loop
			}
		}
	}
}
