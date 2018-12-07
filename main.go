package main

import (
	"fmt"
	"math/rand"

	"github.com/r-tree/model"
)

func main() {
	const COORD_N = 10
	var coordList []model.Coordinate

	// generate random coordinate
	for idx := 0; idx < COORD_N; idx++ {
		coordinate := model.Coordinate{
			X: float64(rand.Intn(20)),
			Y: float64(rand.Intn(20)),
		}
		coordList = append(coordList, coordinate)
	}

	fmt.Print(coordList)
}
