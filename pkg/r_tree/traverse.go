package rtree

import (
	"fmt"

	"github.com/r-tree/pkg/r_tree/model"
)

func (rtree *RTree) Traverse(node model.Node) {
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
			rtree.Traverse(bNode.Entries[idx])
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
