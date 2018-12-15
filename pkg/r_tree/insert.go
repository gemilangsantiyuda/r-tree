package rtree

import (
	"github.com/r-tree/pkg/coordinate"
	"github.com/r-tree/pkg/r_tree/model"
	"github.com/r-tree/pkg/rectangle"
)

// Insert insert an object with index to the tree
func (rtree *RTree) Insert(index int, coordinate *coordinate.Coordinate) {
	rect := rectangle.NewRectangle(coordinate, coordinate)
	leafEntry := &model.LeafEntry{
		Parent:    nil,
		Rectangle: rect,
		Index:     index,
	}

	insertedLeafNode := rtree.ChooseLeaf(rtree.Root, index, rect)
	insertedLeafNode.Insert(leafEntry)

	// if after the insertion the size of entries in insertedLeafNode exceeds max size
	// then invoke split
	if len(insertedLeafNode.Entries) > rtree.MaxEntry {
		insertedNode, newNode := rtree.Split(insertedLeafNode)
		rtree.AdjustTree(insertedNode, newNode)
	} else {
		rtree.AdjustTree(insertedLeafNode, nil)
	}

	// for _, entry := range insertedLeafNode.Entries {
	// 	fmt.Print(entry.Rectangle.LowerLeft)
	// }

}
