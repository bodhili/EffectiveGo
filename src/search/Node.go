package search

import "fmt"

type Node struct {
	Value    int
	Children []*Node
}

func BreadthFirstSearch(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Value)

		for _, child := range node.Children {
			queue = append(queue, child)
		}
	}
}
