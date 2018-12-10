package model

import "github.com/r-tree/pkg/rectangle"

type LeafEntry struct {
	Parent    *LeafNode
	Rectangle *rectangle.Rectangle
	Index     int
}
