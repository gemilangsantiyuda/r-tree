package rtree

import (
	"github.com/r-tree/pkg/r_tree/model"
	"github.com/r-tree/pkg/rectangle"
)

// Search , given a rectange as the boundary, outputs all coordinates' index on coordinate list
// which are located inside of the boundary
func (rtree *RTree) Search(rect *rectangle.Rectangle) []int {
	// traverse(rtree.Root)
	var indexList []int
	indexList = rtree.searchTree(rtree.Root, rect, indexList)
	return indexList
}

func (rtree *RTree) searchTree(node model.Node, rect *rectangle.Rectangle, indexList []int) []int {

	// check if node is leafNode
	if _, ok := node.(*model.LeafNode); ok {
		lNode := node.(*model.LeafNode)
		for idx := range lNode.Entries {
			entry := lNode.Entries[idx]
			// fmt.Println(*entry.Rectangle.LowerLeft)
			if rectangle.IsOverlap(entry.Rectangle, rect) {
				indexList = append(indexList, entry.Index)
			}
		}
		return indexList
	}

	bNode := node.(*model.BranchNode)
	for idx := range bNode.Entries {
		entry := bNode.Entries[idx]
		// fmt.Println(entry.GetRectangle().LowerLeft, entry.GetRectangle().UpperRight)
		if rectangle.IsOverlap(entry.GetRectangle(), rect) {
			indexList = rtree.searchTree(entry, rect, indexList)
		}
	}
	return indexList
}
