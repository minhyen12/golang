package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func testWalk() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 1; i < 11; i++ {
		if i != <-ch {
			fmt.Println("Walk doesn't work")
			break
		}
	}
	fmt.Println("Walk does work")
}
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < cap(ch1); i++ {
		if (<-ch1 != <-ch2) {
			return false
		}
	}
	return true
}
func testSame() {
	if Same(tree.New(1), tree.New(1)) && Same(tree.New(1), tree.New(2)) {
		fmt.Println("Same does work")
	} else {
		fmt.Println("Same doesn't work")
	}
}

func main() {
	testWalk()
	testSame()
}
