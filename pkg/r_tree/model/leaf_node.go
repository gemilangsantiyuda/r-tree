package model

import (
	"github.com/r-tree/pkg/rectangle"
)

type LeafNode struct {
	Rectangle *rectangle.Rectangle
	Entries   []*LeafEntry
}

func (ln *LeafNode) UpdateRectangle() {

	NewRectangle := ln.Entries[0].Rectangle
	for _, entry := range ln.Entries {
		NewRectangle = rectangle.GetCombinedRectangle(NewRectangle, entry.Rectangle)
	}
	ln.Rectangle = NewRectangle
}
