package main

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/r-tree/pkg/coordinate"
	rtree "github.com/r-tree/pkg/r_tree"
)

func main() {
	const coordN = 10
	const minEntry = 2
	const maxEntry = 4
	var coordList []coordinate.Coordinate

	// generate random coordinate
	for idx := 0; idx < coordN; idx++ {
		coordinate := coordinate.Coordinate{
			X: float64(rand.Intn(20)),
			Y: float64(rand.Intn(20)),
		}
		coordList = append(coordList, coordinate)
	}

	Tree := rtree.NewTree(minEntry, maxEntry)

	fmt.Print(reflect.TypeOf(Tree))
}
