package collection

import (
	"fmt"
)

type TreeNode[T any] struct {
	Value T
	left  *TreeNode[T]
	right *TreeNode[T]
}

type Comparator[T any] func(a, b T) int

type Tree[T any] struct {
	Root       *TreeNode[T]
	comparator Comparator[T]
}

func NewTree[T any](cmp Comparator[T]) *Tree[T] {
	return &Tree[T]{comparator: cmp}
}

func (t *Tree[T]) Add(value T) {
	node := &TreeNode[T]{Value: value}

	if t.Root == nil {
		t.Root = node
		return
	}

	curr := t.Root
	for {
		if t.comparator(value, curr.Value) < 0 {
			if curr.left == nil {
				curr.left = node
				return
			}
			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = node
				return
			}
			curr = curr.right
		}
	}
}

func (t *Tree[T]) PrintTree() {
	printPrettyTree(t.Root, "", true)
}

func printPrettyTree[T any](node *TreeNode[T], prefix string, isTail bool) {
	if node == nil {
		return
	}
	fmt.Printf("%s%s%s\n", prefix, ternary(isTail, "└── ", "├── "), fmt.Sprint(node.Value))

	children := []*TreeNode[T]{node.left, node.right}
	n := len(children)
	for i, child := range children {
		if child != nil {
			newPrefix := prefix + ternary(isTail, "    ", "│   ")
			printPrettyTree(child, newPrefix, i == n-1)
		}
	}
}

func ternary(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}

func (t *Tree[T]) InOrder() {
	t.inOrder(t.Root)
	fmt.Println()
}

func (t *Tree[T]) inOrder(node *TreeNode[T]) {
	if node == nil {
		return
	}
	t.inOrder(node.left)
	fmt.Print(node.Value, " ")
	t.inOrder(node.right)
}
