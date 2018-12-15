package model

import (
	"github.com/r-tree/pkg/rectangle"
)

// Node interface
type Node interface {
	GetRectangle() *rectangle.Rectangle
	GetParent() Node
	SetParent(Node)
}
