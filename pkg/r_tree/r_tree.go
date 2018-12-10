package rtree

import "github.com/r-tree/pkg/r_tree/model"

// RTree struct to store r-tree and its constraint
type RTree struct {
	MinEntry, MaxEntry int
	Root               model.Node
}
