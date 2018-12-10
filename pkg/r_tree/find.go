package rtree

import (
	"github.com/r-tree/pkg/r_tree/model"
	"github.com/r-tree/pkg/rectangle"
)

//FindLeaf : Given an index from the list of coordinate, find the leaf node that contains
func (rtree *RTree) FindLeaf(node model.Node, index int, rect *rectangle.Rectangle) *model.LeafNode {

	// check if node is branch node or leaf node (LeafNode)
	if _, ok := node.(*model.LeafNode); ok {
		lNode := node.(*model.LeafNode)
		// Check if index exist in its entries
		for _, entry := range lNode.Entries {
			if entry.Index == index {
				return lNode
			}
		}
		// else return nil
		return nil
	}

	// then node is branch node
	bNode := node.(*model.BranchNode)
	// check all entries that has its rectangles overlaps rect
	for _, entry := range bNode.Entries {
		if rectangle.IsOverlap(entry.GetRectangle(), rect) {
			foundLeaf := rtree.FindLeaf(entry, index, rect)
			if foundLeaf != nil {
				return foundLeaf
			}
		}
	}
	return nil
}
