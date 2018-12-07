package model

import "github.com/r-tree/pkg/rectangle"

type BranchNode struct {
	Rectangle *rectangle.Rectangle
	Entries   []*BranchEntry
}
