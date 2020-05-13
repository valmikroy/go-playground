/*

Lets make it more simple

6 10
8 15
7 19
2 6
15 21
3 9
20 23


This becomes even simple if this list is sorted on the first field.

*/
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := []string{"6 10", "8 15", "7 19", "2 6", "15 21", "3 9", "20 23"}

	var in []int
	var out []int

	for _, data := range input {
		d := strings.Split(data, " ")
		n, _ := strconv.Atoi(d[0])
		in = append(in, n)
		n, _ = strconv.Atoi(d[1])
		out = append(out, n)
	}

	sort.Ints(in)
	sort.Ints(out)

	stop := 16

	total_people_in := 0

	for i := 0; i < len(input); i++ {
		if in[i] <= stop {
			total_people_in += 1
		}
	}

	for i := 0; i < len(input); i++ {
		if out[i] <= stop {
			total_people_in -= 1
		}
	}

	fmt.Println(total_people_in)
}
