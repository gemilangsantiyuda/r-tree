package priorityQueue

import "github.com/r-tree/pkg/r_tree/model"

// Item , data type for element of priority queue
type Item struct {
	Object   model.Node
	Priority float64
	Index    int
}
