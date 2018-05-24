package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

var count int = 1

func Walk(t *tree.Tree, ch chan int){
	WalkRecursive(t , ch )
	close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkRecursive(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		WalkRecursive(t.Right,ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for k := range c1{
		v1 := k
		v2 := <- c2
		if v1!=v2 {return false}
	}
	return true
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	t1 := tree.New(1)
	t2 := tree.New(1)
	go Walk(t1, c1)

	for i := range c1{
		fmt.Print(i, " ")
	}

	fmt.Println()

	go Walk(t2, c2)

	for i := range c2{
		fmt.Print(i, " ")
	}

	fmt.Println()
	fmt.Print(Same(t1, t2))
}}