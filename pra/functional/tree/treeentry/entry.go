package main

import (
	"fmt"
	"pra/functional/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = new(tree.Node)
	root.Right.Left.SetValue(4)

	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
}
