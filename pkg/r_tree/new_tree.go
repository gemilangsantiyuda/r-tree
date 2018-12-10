package r_tree

import (
	"github.com/r-tree/pkg/r_tree/model"
)

func NewTree(MinEntry, MaxEntry int) *model.RTree {

	newRoot := &model.LeafNode{}

	return &model.RTree{
		MinEntry: MinEntry,
		MaxEntry: MaxEntry,
		Root:     newRoot,
	}
}
