package rectangle

import (
	"math"

	"github.com/r-tree/pkg/coordinate"
)

func NewRectangle(LowerLeft, UpperRight *coordinate.Coordinate) *Rectangle {
	Rect := &Rectangle{
		LowerLeft:  LowerLeft,
		UpperRight: UpperRight,
	}
	Rect.Area = ((UpperRight.Y - LowerLeft.Y) * (UpperRight.X - LowerLeft.X))
	return Rect
}

func GetCombinedRectangle(Rect1, Rect2 *Rectangle) *Rectangle {
	UpperRight := coordinate.Coordinate{
		X: math.Max(Rect1.UpperRight.X, Rect2.UpperRight.X),
		Y: math.Max(Rect1.UpperRight.Y, Rect2.UpperRight.Y),
	}
	LowerLeft := coordinate.Coordinate{
		X: math.Max(Rect1.LowerLeft.X, Rect2.LowerLeft.X),
		Y: math.Max(Rect1.LowerLeft.Y, Rect2.LowerLeft.Y),
	}

	return NewRectangle(LowerLeft, UpperRight)
}
