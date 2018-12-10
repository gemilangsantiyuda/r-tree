package model

import "github.com/r-tree/pkg/rectangle"

// LeafEntry struct for the entries of the leaf node
type LeafEntry struct {
	Parent    *LeafNode
	Rectangle *rectangle.Rectangle
	Index     int
}
