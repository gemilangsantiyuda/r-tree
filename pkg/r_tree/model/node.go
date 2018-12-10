package model

import (
	"github.com/r-tree/pkg/rectangle"
)

type Node interface {
	UpdateRectangle()
	GetRectangle() *rectangle.Rectangle
}
