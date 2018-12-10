package model

import (
	"github.com/r-tree/pkg/rectangle"
)

type BranchNode struct {
	Parent    *BranchNode
	Rectangle *rectangle.Rectangle
	Entries   []Node
}

func (self *BranchNode) GetRectangle() *rectangle.Rectangle {
	return self.Rectangle
}

func (self *BranchNode) UpdateRectangle() {
	// if it has no entries, just return
	if len(self.Entries) == 0 {
		return
	}

	// else create new rectangle out of combined rectangles from all its entries
	var newRect *rectangle.Rectangle
	*newRect = *self.Entries[0].GetRectangle()

	for _, entry := range self.Entries {
		rect := entry.GetRectangle()
		newRect = rectangle.GetCombinedRectangle(newRect, rect)
	}

	self.Rectangle = newRect
}
