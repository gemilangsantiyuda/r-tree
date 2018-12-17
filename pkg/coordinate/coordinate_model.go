package coordinate

//Coordinate struct for coordinate, thats all
type Coordinate struct {
	X, Y float64
}

// GetDistance get distance from this coordinate to other coordinate or to rectangle
func (coord *Coordinate) GetDistance(object interface{}) float64 {

	if _, ok := object.(*Coordinate); ok {
		coord2 := object.(*Coordinate)
		return coord.X-
	}
}
