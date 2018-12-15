package rectangle

import (
	"github.com/r-tree/pkg/coordinate"
)

//Rectangle bla bla bla
type Rectangle struct {
	LowerLeft, UpperRight *coordinate.Coordinate
	Area                  float64
}

// // ContainCoordinate if this rect is the boundary of certain coordinate
// func (rect *Rectangle) ContainCoordinate(coordinate *coordinate.Coordinate) bool {

// 	coordinateInside := ((coordinate.X >= rect.LowerLeft.X) && (coordinate.X <= rect.UpperRight.X) && (coordinate.Y >= rect.LowerLeft.Y) && (coordinate.Y <= rect.UpperRight.Y))

// 	return coordinateInside
// }
