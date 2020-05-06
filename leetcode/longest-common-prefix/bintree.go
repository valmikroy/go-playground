package main

import "fmt"

type Tree struct {
	Index  int
	Weight int
	Char   string
	Left   *Tree
	Right  *Tree
}

func NewTree(c string) *Tree {

	return &Tree{
		Index:  0,
		Weight: 1,
		Char:   c,
		Left:   nil,
		Right:  nil,
	}
}

func (t *Tree) Insert(idx int, c string) {

	switch {
	case t.Index > idx:
		if t.Left == nil {
			t.Left = &Tree{
				Index:  idx,
				Weight: 1,
				Char:   c,
				Left:   nil,
				Right:  nil,
			}
		} else {
			t.Left.Insert(idx, c)
		}
	case t.Index < idx:
		if t.Right == nil {
			t.Right = &Tree{
				Index:  idx,
				Weight: 1,
				Char:   c,
				Left:   nil,
				Right:  nil,
			}
		} else {
			t.Right.Insert(idx, c)
		}
	case t.Index == idx:
		if t.Char == c {
			t.Weight += 1
		}
	}
}

func (t *Tree) Traverse() {

	if t.Left != nil {
		t.Left.Traverse()
	}

	fmt.Printf("%d - %s - %d\n", t.Index, t.Char, t.Weight)

	if t.Right != nil {
		t.Right.Traverse()
	}

}

func main() {

	var x string
	x = "geeksforgeeks"

	var t *Tree

	for i := 0; i < len(x); i++ {
		if i == 0 {
			t = NewTree(string(x[i]))
		} else {
			t.Insert(i, string(x[i]))
		}
	}

	x = "geeks"
	// input := []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	for i := 0; i < len(x); i++ {
		t.Insert(i, string(x[i]))
	}

	x = "geek"
	for i := 0; i < len(x); i++ {
		t.Insert(i, string(x[i]))
	}

	x = "geezer"
	for i := 0; i < len(x); i++ {
		t.Insert(i, string(x[i]))
	}

	t.Traverse()

}
