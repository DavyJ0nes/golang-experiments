package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	nodes := Read()
	GenerateGraphViz(nodes)
}

// Read is used to traverse the input file to build the tree
func Read() []Node {
	var N int
	fmt.Scanf("%d", &N)

	var nodes = make([]Node, N)

	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		nodes[i].Value = val

		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}

		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
	}

	return nodes
}

// PrintNode is a helper function to inspect the values of a tree
func PrintNode(n *Node) {
	leftVal := -1
	rightVal := -1

	if n.Left != nil {
		leftVal = n.Left.Value
	}

	if n.Right != nil {
		rightVal = n.Right.Value
	}

	fmt.Printf("Val: %d \t| Left: %d \t| Right: %d\n", n.Value, leftVal, rightVal)
}

// GenerateGraphViz converts the tree into correct format for graphviz to read
func GenerateGraphViz(nodes []Node) {
	var fileLines []string
	fileLines = append(fileLines, "digraph TestGraph {")

	for _, node := range nodes {
		fileLines = append(fileLines, fmt.Sprintf("\t%d", node.Value))

		if node.Left != nil {
			line := fmt.Sprintf("\t%d -> %d", node.Value, node.Left.Value)
			fileLines = append(fileLines, line)
		}

		if node.Right != nil {
			line := fmt.Sprintf("\t%d -> %d", node.Value, node.Right.Value)
			fileLines = append(fileLines, line)
		}
	}

	fileLines = append(fileLines, "}")

	for _, line := range fileLines {
		fmt.Println(line)
	}
}
