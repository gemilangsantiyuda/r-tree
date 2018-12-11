package model

import "github.com/r-tree/pkg/rectangle"

// LeafEntry struct for the entries of the leaf node
type LeafEntry struct {
	Parent    *LeafNode
	Rectangle *rectangle.Rectangle
	Index     int
}

// UpdateRectangle method so that this struct implements Node
func (le *LeafEntry) UpdateRectangle() {
	// update rectangle unnecessary
	return
}

// GetRectangle method to return its rectangle (implement Node)
func (le *LeafEntry) GetRectangle() *rectangle.Rectangle {
	return le.Rectangle
}

// GetParent method to implement Node
func (le *LeafEntry) GetParent() Node {
	return le.Parent
}

// SetParent method to implement Node
func (le *LeafEntry) SetParent(ln Node) {
	le.Parent = ln.(*LeafNode)
}
