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

/*
NewTree creates and returns a new instance of Tree using the provided comparator function.
The comparator function determines the order of elements within the tree.
*/
func NewTree[T any](cmp Comparator[T]) *Tree[T] {
	return &Tree[T]{comparator: cmp}
}

/* Add adds a new node with the given value to the tree, keeping the tree balanced. */
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

/* PrintTree prints the tree in a visually appealing format. */
func (t *Tree[T]) PrintTree() {
	printPrettyTree(t.Root, "", true)
}

/*
printPrettyTree prints a tree in a visually appealing format, with each node
indented underneath its parent and connected by lines. The isTail parameter
indicates whether the node is the last child of its parent, and is used to
determine which type of line to draw.
*/
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

/* InOrder performs an in-order traversal of the tree and prints each node's value. */
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
