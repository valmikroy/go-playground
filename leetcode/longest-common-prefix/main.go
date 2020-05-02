/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".


Input: ["flower","flow","flight"]
Output: "fl"

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

*/
package main

import "fmt"

type Data struct {
	r string
}

func NewData(input string) *Data {
	return &Data{
		r: input,
	}
}

func (h *Data) hello() string {
	return h.r
}
