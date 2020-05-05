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

func main() {

	//input := []string{"flower", "flow", "flight"}
	input := []string{"geeksforgeeks", "geeks", "geek", "geezer"}

	withChar(input)
}

// with regex
func withChar(input []string) {

	longest := 0

	prefix := ""

	for _, c := range input {
		if len(c) > longest {
			longest = len(c)
		}
	}

	for i := 0; i < longest; i++ {

		letter := ""

		for j, c := range input {

			if i > len(c)-1 {
				break
			}

			char := string(c[i])
			if j == 0 {
				letter = char
			}

			if letter != char {
				break
			}

			if j == len(input)-1 {
				prefix = prefix + letter
			}

		}
	}

	fmt.Println(prefix)
}
