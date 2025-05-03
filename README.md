# Collections

This repository provides a collection of data structures implemented in Go, including Stack, Queue, LinkedList, Set, and Tree. These data structures are generic and can be used with any data type.

## Getting Started

To use these data structures in your Go project, you can import the desired package and create instances of the data structures as needed.

### Prerequisites

- Go programming language (version 1.18 or later)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/collections.git
   ```

2. Navigate to the project directory:

   ```bash
   cd collections
   ```

## Usage

### Stack

```go
stack := collection.NewStack[int]()
stack.Push(1)
stack.Push(2)
fmt.Println(stack.Pop()) // Output: 2
```

### Queue

```go
queue := collection.NewQueue[int]()
queue.Enqueue(1)
queue.Enqueue(2)
fmt.Println(queue.Dequeue()) // Output: 1
```

### LinkedList

```go
list := collection.NewLinkedList[int]()
list.Add(1)
list.Add(2)
list.Print() // Output: 1 -> 2
```

### Set

```go
set := collection.NewSet[int]()
set.Add(1)
set.Add(2)
fmt.Println(set.Contains(1)) // Output: true
```

### Tree

```go
tree := collection.NewTree[int](func(a, b int) int { return a - b })
tree.Add(1)
tree.Add(2)
tree.PrintTree()
```
