package rtree

import "github.com/r-tree/pkg/r_tree/model"

// CondenseTree to condense the tree, starting from a leafnode that has just removed one of its entries
func (rtree *RTree) CondenseTree(leafNode *model.LeafNode) {

	var removedEntries []*model.LeafEntry
	removedEntries = rtree.condense(leafNode, removedEntries)
	for idx := range removedEntries {
		entry := removedEntries[idx]
		rtree.Insert(entry.Index, entry.Rectangle.LowerLeft)
	}
}

func (rtree *RTree) condense(node model.Node, removedEntries []*model.LeafEntry) []*model.LeafEntry {

	// if root is reached, then the condense process stopped, returning all the removed entries to be reinserted
	if node == rtree.Root {
		return removedEntries
	}

	// if it's leaf node
	if _, ok := node.(*model.LeafNode); ok {
		lNode := node.(*model.LeafNode)
		parent := lNode.Parent
		// if the size of entries is smaller than the tree's minimum, then delete this node
		// but save all its entries first
		if len(lNode.Entries) < rtree.MinEntry {
			removedEntries = rtree.getEntries(lNode, removedEntries)
			parent.RemoveChild(lNode)
		}
		removedEntries = rtree.condense(parent, removedEntries)
	} else {
		bNode := node.(*model.BranchNode)
		parent := bNode.Parent
		// if the size of entries is smaller than the tree's minimum, then delete this node
		// but save all its entries first
		if len(bNode.Entries) < rtree.MinEntry {
			removedEntries = rtree.getEntries(bNode, removedEntries)
			parent.RemoveChild(bNode)
		}
		removedEntries = rtree.condense(parent, removedEntries)
	}
	return removedEntries
}

func (rtree *RTree) getEntries(node model.Node, entries []*model.LeafEntry) []*model.LeafEntry {

	// if this is leaf node then save all its entries into -> entries
	if _, ok := node.(*model.LeafNode); ok {
		lNode := node.(*model.LeafNode)
		for idx := range lNode.Entries {
			entry := lNode.Entries[idx]
			entries = append(entries, entry)
		}
		return entries
	}

	// else traverse all its child and save their entries
	bNode := node.(*model.BranchNode)
	for idx := range bNode.Entries {
		entry := bNode.Entries[idx]
		entries = rtree.getEntries(entry, entries)
	}
	return entries
}
