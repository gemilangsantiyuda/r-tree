package model

import (
	"github.com/r-tree/pkg/rectangle"
)

// LeafNode struct
type LeafNode struct {
	Parent    *BranchNode
	Rectangle *rectangle.Rectangle
	Entries   []*LeafEntry
}

// GetRectangle fucntion so that it implements Node
func (ln *LeafNode) GetRectangle() *rectangle.Rectangle {
	return ln.Rectangle
}

// UpdateRectangle function to update rectangle as the combined areas of its entries' rectangle
func (ln *LeafNode) UpdateRectangle() {
	// if it has no entries, just return
	if len(ln.Entries) == 0 {
		return
	}

	// else create new rectangle out of combined rectangles from all its entries
	newRect := &rectangle.Rectangle{}
	*newRect = *ln.Entries[0].Rectangle

	for _, entry := range ln.Entries {
		rect := entry.Rectangle
		newRect = rectangle.GetCombinedRectangle(newRect, rect)
	}

	ln.Rectangle = newRect
}

// Insert insert a leaf entry into the leaf node's entries
func (ln *LeafNode) Insert(leafEntry *LeafEntry) {
	leafEntry.Parent = ln
	ln.Entries = append(ln.Entries, leafEntry)
}

// GetParent get this node's parent node
func (ln *LeafNode) GetParent() Node {
	return ln.Parent
}

// SetParent set this node's parent node
func (ln *LeafNode) SetParent(node Node) {
	ln.Parent = node.(*BranchNode)
}
