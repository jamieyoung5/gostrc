package dlx

import (
	"fmt"
)

type Matrix struct {
	Root *Node
}

type Node struct {
	Left, Right, Up, Down *Node
	Column                *Column
}

// NewMatrix generates columns and each columns' respective header node
func NewMatrix(columnIDs []string) *Matrix {
	root := &Column{Node: &Node{}}
	root.Left = root.Node
	root.Right = root.Node

	lastHeader := root
	for _, id := range columnIDs {
		newHeader := &Column{ID: id}
		newHeader.Node = &Node{Left: lastHeader.Node, Right: root.Node, Column: newHeader}
		newHeader.Up = newHeader.Node
		newHeader.Down = newHeader.Node

		lastHeader.Right = newHeader.Node
		root.Left = newHeader.Node
		lastHeader = newHeader
	}

	return &Matrix{Root: root.Node}
}

func (m *Matrix) PrintMatrix() {
	currentHeader := m.Root.Right // Start from the first actual header, not the root itself
	for currentHeader != m.Root {
		fmt.Printf("Column ID: %s, Size: %d\n", currentHeader.Column.ID, currentHeader.Column.Size)
		currentNode := currentHeader.Down
		for currentNode != currentHeader { // Iterate until we loop back to the dummy node
			fmt.Println("  Node at", currentNode)
			currentNode = currentNode.Down
		}
		currentHeader = currentHeader.Right
	}
}

// AppendRow adds a new node to the bottom of every column included in columnIds and links them together to form a row
func (m *Matrix) AppendRow(columnIds []string) {
	idSet := make(map[string]struct{})
	for _, id := range columnIds {
		idSet[id] = struct{}{}
	}

	var firstNode, lastNode *Node

	header := m.Root.Right
	for header != m.Root {
		if _, exists := idSet[header.Column.ID]; exists {
			newNode := &Node{
				Column: header.Column,
				Up:     header.Up,
				Down:   header,
			}

			// Link vertically
			newNode.Up.Down = newNode
			newNode.Down.Up = newNode
			header.Column.Size++

			// Link horizontally
			if firstNode == nil {
				firstNode = newNode
				lastNode = newNode
			} else {
				newNode.Left = lastNode
				lastNode.Right = newNode
				lastNode = newNode
			}
		}
		header = header.Right
	}

	// Close the horizontal loop
	if firstNode != nil {
		if lastNode == firstNode {
			// If there is only one node, it should point to itself
			firstNode.Left = firstNode
			firstNode.Right = firstNode
		} else {
			firstNode.Left = lastNode
			lastNode.Right = firstNode
		}
	}
}

func (m *Matrix) Search(solution []*Node, solutions *[][]*Node, maxSolutions int) {
	if len(*solutions) >= maxSolutions {
		return
	}

	// Check if there are no more columns to cover
	if m.Root.Right == m.Root {
		// All columns are covered, hence a solution is found
		// Copy the solution to avoid referencing the same slice
		solved := make([]*Node, len(solution))
		copy(solved, solution)
		*solutions = append(*solutions, solved)
		return
	}

	// Choose the column with the fewest nodes to minimize the branching factor
	c := m.getColumnWithFewestNodes()

	c.Cover()

	for r := c.Down; r != c.Node; r = r.Down {
		solution = append(solution, r)

		// Cover all columns in this row
		for j := r.Right; j != r; j = j.Right {
			j.Column.Cover()
		}

		m.Search(solution, solutions, maxSolutions)

		// Uncover all columns in this row for backtracking
		for j := r.Left; j != r; j = j.Left {
			j.Column.Uncover()
		}

		// Remove the row from the solution set
		solution = solution[:len(solution)-1]
	}

	c.Uncover()
}

// getColumnWithFewestNodes get the smallest column in the matrix
func (m *Matrix) getColumnWithFewestNodes() *Column {
	minSize := int(^uint(0) >> 1) // Max int value
	chosenColumn := m.Root.Right.Column
	for c := m.Root.Right; c != m.Root; c = c.Right {
		if c.Column.Size < minSize {
			minSize = c.Column.Size
			chosenColumn = c.Column
		}
	}
	return chosenColumn
}
