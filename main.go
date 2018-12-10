package main

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/r-tree/pkg/coordinate"
	rtree "github.com/r-tree/pkg/r_tree"
)

func main() {
	const COORD_N = 10
	const MIN_ENTRY = 2
	const MAX_ENTRY = 4
	var coordList []coordinate.Coordinate

	// generate random coordinate
	for idx := 0; idx < COORD_N; idx++ {
		coordinate := coordinate.Coordinate{
			X: float64(rand.Intn(20)),
			Y: float64(rand.Intn(20)),
		}
		coordList = append(coordList, coordinate)
	}

	Tree := rtree.NewTree(MIN_ENTRY, MAX_ENTRY)

	fmt.Print(reflect.TypeOf(Tree))
}
