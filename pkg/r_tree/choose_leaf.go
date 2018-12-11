package rtree

import (
	"math"

	"github.com/r-tree/pkg/r_tree/model"
	"github.com/r-tree/pkg/rectangle"
)

// ChooseLeaf choose a leaf node to insert object of index = index
func (rtree *RTree) ChooseLeaf(node model.Node, index int, rect *rectangle.Rectangle) *model.LeafNode {

	// check if the node is leaf, then return the node
	if _, ok := node.(*model.LeafNode); ok {
		return node.(*model.LeafNode)
	}

	// else if branch node, then find the node that
	// has new rectangle with the smallest area
	var bestNode model.Node
	minimumArea := math.Inf(1)
	// cast to branch node
	branchNode := node.(*model.BranchNode)
	for _, entry := range branchNode.Entries {
		newRect := rectangle.GetCombinedRectangle(rect, entry.GetRectangle())
		if newRect.Area < minimumArea {
			minimumArea = newRect.Area
			bestNode = entry
		}
	}
	return rtree.ChooseLeaf(bestNode, index, rect)
}
