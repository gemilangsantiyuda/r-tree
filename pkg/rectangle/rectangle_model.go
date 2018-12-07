package rectangle

import (
	"github.com/r-tree/pkg/coordinate"
)

//Rectangle bla bla bla
type Rectangle struct {
	LowerLeft, UpperRight *coordinate.Coordinate
	Area                  float64
}
