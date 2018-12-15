package main

import (
	"fmt"
	"math/rand"

	"github.com/r-tree/pkg/rectangle"

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

	rect := &rectangle.Rectangle{
		LowerLeft: &coordinate.Coordinate{
			X: 14,
			Y: 0,
		},
		UpperRight: &coordinate.Coordinate{
			X: 17,
			Y: 17,
		},
	}
	indexes := Tree.Search(rect)
	fmt.Println("___")
	for _, index := range indexes {
		fmt.Println(coordList[index])
	}
}
