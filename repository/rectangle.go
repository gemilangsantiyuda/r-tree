package repository

import "github.com/r-tree/model"

func NewRectangle(LowerLeft, UpperRight model.Coordinate) *model.Rectangle {
	Rect := &model.Rectangle{
		LowerLeft:  LowerLeft,
		UpperRight: UpperRight,
	}
	Rect.Area = ((UpperRight.Y - LowerLeft.Y) * (UpperRight.X - LowerLeft.X))
	return Rect
}

func GetCombinedRectangle(Rect1, Rect2 model.Rectangle) *model.Rectangle {
	UpperRight := Coordinate{
		X: max(Rect1.UpperRight.X, Rect2.UpperRight.X),
		Y: max(Rect1.UpperRight.Y, Rect2.UpperRight.Y),
	}
	LowerLeft := Coordinate{
		X: max(Rect1.LowerLeft.X, Rect2.LowerLeft.X),
		Y: max(Rect1.LowerLeft.Y, Rect2.LowerLeft.Y),
	}

	return NewRectangle(LowerLeft, UpperRight)
}
