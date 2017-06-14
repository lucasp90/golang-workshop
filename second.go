package main

import (
	"fmt"
)

type tree struct {
	value int
	left *tree
	right *tree
}

func add(t *tree, value int) *tree {
	if t==nil {
		t = new(tree)
		t.value = value
		return t
	}

	if t.value >= value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}


func (t *tree) String() string {
	if t==nil {
		return ""
	}
	if t.left == nil && t.right == nil {
		return fmt.Sprintf("%d", t.value)
	}
	if t.left == nil {
		return fmt.Sprintf("%d,%s", t.value, t.right.String())
	} else if t.right == nil {
		return fmt.Sprintf("%s,%d", t.left.String(), t.value)
	} else {
		return fmt.Sprintf("%s,%d,%s", t.left.String(), t.value, t.right.String())
	}

}

func main() {
	var t = new(tree)
	numbers := [10]int{1,3,2,4,5,6,21,44,8,15}

	for _, n := range numbers {
		t = add(t, n)
	}

	fmt.Println(t.String())

}