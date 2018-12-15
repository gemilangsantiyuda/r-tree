package main

import (
	"fmt"
	"math/rand"

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
	for _, coord := range coordList {
		fmt.Println(coord)
	}
	for idx := range coordList {
		coord := coordList[idx]
		Tree.Insert(idx, &coord)
		// fmt.Println("success", idx+1)
	}
}
