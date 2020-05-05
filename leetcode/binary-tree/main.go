package main

import "github.com/valmikroy/go-bintree/pkg/bintree"

func main() {

	x := bintree.NewSimpleTree(10)

	x.Insert(1)
	x.Insert(10)
	x.Insert(9)
	x.Insert(13)
	x.Insert(11)
	x.Insert(12)

	x.Traverse()

}
