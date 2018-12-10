package model

import (
	"github.com/r-tree/pkg/rectangle"
)

type LeafNode struct {
	Parent    *BranchNode
	Rectangle *rectangle.Rectangle
	Entries   []*LeafEntry
}

func (self *LeafNode) GetRectangle() *rectangle.Rectangle {
	return self.Rectangle
}

func (self *LeafNode) UpdateRectangle() {
	// if it has no entries, just return
	if len(self.Entries) == 0 {
		return
	}

	// else create new rectangle out of combined rectangles from all its entries
	var newRect *rectangle.Rectangle
	*newRect = *self.Entries[0].Rectangle

	for _, entry := range self.Entries {
		rect := entry.Rectangle
		newRect = rectangle.GetCombinedRectangle(newRect, rect)
	}

	self.Rectangle = newRect
}
