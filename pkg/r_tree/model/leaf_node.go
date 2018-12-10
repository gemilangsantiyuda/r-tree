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
	var newRect *rectangle.Rectangle
	*newRect = *ln.Entries[0].Rectangle

	for _, entry := range ln.Entries {
		rect := entry.Rectangle
		newRect = rectangle.GetCombinedRectangle(newRect, rect)
	}

	ln.Rectangle = newRect
}
