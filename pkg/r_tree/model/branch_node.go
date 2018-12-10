package model

import (
	"github.com/r-tree/pkg/rectangle"
)

// BranchNode struct for branch nodes
type BranchNode struct {
	Parent    *BranchNode
	Rectangle *rectangle.Rectangle
	Entries   []Node
}

// GetRectangle function so it implements Node
func (bn *BranchNode) GetRectangle() *rectangle.Rectangle {
	return bn.Rectangle
}

// UpdateRectangle updating rectangle as combined area of its entries' rectangles
func (bn *BranchNode) UpdateRectangle() {
	// if it has no entries, just return
	if len(bn.Entries) == 0 {
		return
	}

	// else create new rectangle out of combined rectangles from all its entries
	var newRect *rectangle.Rectangle
	*newRect = *bn.Entries[0].GetRectangle()

	for _, entry := range bn.Entries {
		rect := entry.GetRectangle()
		newRect = rectangle.GetCombinedRectangle(newRect, rect)
	}

	bn.Rectangle = newRect
}
