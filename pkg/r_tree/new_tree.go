package rtree

import (
	"github.com/r-tree/pkg/r_tree/model"
)

// NewTree function to init new r-tree
func NewTree(MinEntry, MaxEntry int) *RTree {

	newRoot := &model.LeafNode{}

	return &RTree{
		MinEntry: MinEntry,
		MaxEntry: MaxEntry,
		Root:     newRoot,
	}
}
