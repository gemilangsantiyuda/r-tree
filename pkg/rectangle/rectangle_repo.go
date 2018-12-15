package rectangle

import (
	"math"

	"github.com/r-tree/pkg/coordinate"
)

func NewRectangle(lowerLeft, upperRight *coordinate.Coordinate) *Rectangle {
	rect := &Rectangle{
		LowerLeft:  lowerLeft,
		UpperRight: upperRight,
	}
	rect.Area = ((upperRight.Y - lowerLeft.Y) * (upperRight.X - lowerLeft.X))
	return rect
}

func GetCombinedRectangle(rect1, rect2 *Rectangle) *Rectangle {
	UpperRight := &coordinate.Coordinate{
		X: math.Max(rect1.UpperRight.X, rect2.UpperRight.X),
		Y: math.Max(rect1.UpperRight.Y, rect2.UpperRight.Y),
	}
	LowerLeft := &coordinate.Coordinate{
		X: math.Min(rect1.LowerLeft.X, rect2.LowerLeft.X),
		Y: math.Min(rect1.LowerLeft.Y, rect2.LowerLeft.Y),
	}

	return NewRectangle(LowerLeft, UpperRight)
}

// IsOverlap : check if 2 rectangles overlap, if a rect1 overlaps rect2 then either rect1 lowerleft coordinate is inside rect2 or rect1 upperright coordinate is inside rect2
func IsOverlap(rect1, rect2 *Rectangle) bool {

	// if rect1 is above rect2
	if rect1.LowerLeft.Y > rect2.UpperRight.Y {
		return false
	}
	// if rect2 is above rect1
	if rect2.LowerLeft.Y > rect1.UpperRight.Y {
		return false
	}
	// if rect1 is on the left of rect2
	if rect1.UpperRight.X < rect2.LowerLeft.X {
		return false
	}
	// if rect2 is on the left of rect1
	if rect2.UpperRight.X < rect1.LowerLeft.X {
		return false
	}

	return true
}
