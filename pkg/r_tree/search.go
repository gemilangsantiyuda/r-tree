package rtree

import (
	"fmt"

	"github.com/r-tree/pkg/r_tree/model"
	"github.com/r-tree/pkg/rectangle"
)

// Search , given a rectange as the boundary, output all coordinates' index on coordinate list
// which are located inside of the boundary
func (rtree *RTree) Search(rect *rectangle.Rectangle) []int {
	// traverse(rtree.Root)
	var indexList []int
	indexList = rtree.searchTree(rtree.Root, rect, indexList)
	return indexList
}

func traverse(node model.Node) {
	fmt.Println("Parent", node, ":")
	rect := node.GetRectangle()
	fmt.Println("rect", rect.LowerLeft, rect.UpperRight)
	if _, ok := node.(*model.BranchNode); ok {
		bNode := node.(*model.BranchNode)
		fmt.Println("CHILD")
		for idx := range bNode.Entries {
			fmt.Println(bNode.Entries[idx])
		}
		fmt.Println("CHILD!")
		for idx := range bNode.Entries {
			traverse(bNode.Entries[idx])
		}
	} else {
		lNode := node.(*model.LeafNode)
		fmt.Println("Entry")
		for idx := range lNode.Entries {
			fmt.Println(lNode.Entries[idx])
		}
		fmt.Println("Entry!")
	}
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
