package distance

import (
	"math"

	"github.com/r-tree/pkg/coordinate"
	"github.com/r-tree/pkg/rectangle"
)

//GetCoordToCoordDistance just calculate haversine distance of 2 given coordinates
func GetCoordToCoordDistance(coordOrigin, coordDestination *coordinate.Coordinate) float64 {
	dx := coordOrigin.X - coordDestination.X
	dy := coordOrigin.Y - coordDestination.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// GetCoordToRectDistance get distance between coordinate and rectangle
// formula is from Russopolous, Kelley, Vincent MINDIST
func GetCoordToRectDistance(coord *coordinate.Coordinate, rect *rectangle.Rectangle) float64 {

	rectX, rectY := coord.X, coord.Y
	if coord.X < rect.LowerLeft.X {
		rectX = rect.LowerLeft.X
	} else if coord.X > rect.UpperRight.X {
		rectX = rect.UpperRight.X
	}

	if coord.Y < rect.LowerLeft.Y {
		rectY = rect.LowerLeft.Y
	} else if coord.Y > rect.UpperRight.Y {
		rectY = rect.UpperRight.Y
	}

	dx := coord.X - rectX
	dy := coord.Y - rectY
	return math.Sqrt(dx*dx + dy*dy)
}
