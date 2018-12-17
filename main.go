package main

import (
	"fmt"
	"math/rand"

	"github.com/r-tree/pkg/coordinate"
	rtree "github.com/r-tree/pkg/r_tree"
)

func main() {
	const coordN = 10000
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

	// rect := &rectangle.Rectangle{
	// 	LowerLeft: &coordinate.Coordinate{
	// 		X: 0,
	// 		Y: 0,
	// 	},
	// 	UpperRight: &coordinate.Coordinate{
	// 		X: 50,
	// 		Y: 50,
	// 	},
	// }
	// Tree.Traverse(Tree.Root)
	// Tree.Remove(1, &coordList[1])
	// fmt.Println("++++++++++++++++++++++++")
	// Tree.Traverse(Tree.Root)
}

func printCoords(indexList []int, coordList []coordinate.Coordinate) {
	fmt.Println("___")
	for _, idx := range indexList {
		fmt.Println(coordList[idx])
	}
}
