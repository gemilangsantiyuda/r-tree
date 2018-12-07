package model

import "github.com/r-tree/pkg/rectangle"

type LeafEntry struct {
	Rectangle *rectangle.Rectangle
	Index     int
}
