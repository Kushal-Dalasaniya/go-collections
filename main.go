package main

import (
	"collections/collection"
	"fmt"
)

func main() {
	stack := collection.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop()) // Output: 2

	queue := collection.NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	fmt.Println(queue.Dequeue()) // Output: 1

	list := collection.NewLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Print() // Output: 1 -> 2

	set := collection.NewSet[int]()
	set.Add(1)
	set.Add(2)
	fmt.Println(set.Contains(1)) // Output: true

	tree := collection.NewTree[int](
		func(a, b int) int {
			return a - b
		})

	tree.Add(4)
	tree.Add(2)
	tree.Add(6)
	tree.Add(1)
	tree.Add(3)
	tree.Add(5)
	tree.Add(8)
	tree.Add(53)
	tree.Add(52)
	tree.Add(54)
	tree.Add(42)

	tree.PrintTree()
}
