/*


Given n non-negative integers a1, a2, ..., an , where each represents a point
at coordinate (i, ai). n vertical lines are drawn such that the two endpoints
of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis
forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.


Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49

*/

package main

import "fmt"

type Tree struct {
	product int
	x       int
	y       int
	left    *Tree
	right   *Tree
}

func (t *Tree) Insert(x int, y int) {

	product := x * y
	switch {
	case product > t.product:
		if t.right == nil {
			t.right = &Tree{
				product: product,
				x:       x,
				y:       y,
				left:    nil,
				right:   nil,
			}
		} else {
			t.right.Insert(x, y)
		}
	case product < t.product:
		if t.left == nil {
			t.left = &Tree{
				product: product,
				x:       x,
				y:       y,
				left:    nil,
				right:   nil,
			}
		} else {
			t.left.Insert(x, y)
		}
	}
}

func (t *Tree) Traverse() {

	if t.right != nil {
		t.right.Traverse()
	}

	fmt.Printf("%d x %d = %d\n", t.x, t.y, t.product)

	if t.left != nil {
		t.left.Traverse()
	}
}

func find_min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	var t *Tree

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {

			var min, diff int
			diff = j - i
			min = find_min(input[i], input[j])

			if t == nil {
				t = &Tree{
					product: min * diff,
					x:       min,
					y:       diff,
					left:    nil,
					right:   nil,
				}
			}

			t.Insert(min, diff)
		}
	}

	t.Traverse()
}
