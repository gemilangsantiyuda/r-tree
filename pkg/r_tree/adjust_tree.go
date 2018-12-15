package rtree

import "github.com/r-tree/pkg/r_tree/model"

// AdjustTree method to adjust tree after inserting new leaf entry, to adjust the new area of nodes and especially to handle splitting that causes another splitting
func (rtree *RTree) adjustTree(node, node2 model.Node) {
	// check if node is the tree root
	if node == rtree.Root {
		// check if node2 is not null or the root was split before
		if node2 != nil {
			// make a new root and node + node2 as its new entries
			newRoot := &model.BranchNode{}
			newRoot.Insert(node)
			newRoot.Insert(node2)
			rtree.Root = newRoot
			return
		}
		// if root was not split then tree adjustment is done
		return
	}

	// if node is not root
	// first get the parent of node
	parent := node.GetParent().(*model.BranchNode)

	// if node was split before (node2 is not nil) then insert it to the parent
	if node2 != nil {
		parent.Insert(node2)
	}
	// if the parent entries' size exceed max entry then split the parent
	if len(parent.Entries) > rtree.MaxEntry {
		parent, parent2 := rtree.split(parent)
		// then adjust tree right up with parent and parent2
		rtree.adjustTree(parent, parent2)
	} else {
		// update parent rectangle and adjust tree up to parent
		rtree.adjustTree(parent, nil)
	}
}
