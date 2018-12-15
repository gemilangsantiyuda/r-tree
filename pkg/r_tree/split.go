package rtree

import (
	"math"

	"github.com/r-tree/pkg/rectangle"

	"github.com/r-tree/pkg/r_tree/model"
)

// Split : in case of exceeding entries' size of a node, then the node has to be split
// the split is of quadratic split
func (rtree *RTree) Split(node model.Node) (model.Node, model.Node) {
	// if node is leafnode then split into two leaf nodes
	if _, ok := node.(*model.LeafNode); ok {
		return splitLeaf(node)
	}
	// else then it is branch node
	return splitBranch(node)
}

func splitBranch(node model.Node) (model.Node, model.Node) {
	bNode := node.(*model.BranchNode)
	// save its entries into another list
	entries := make([]model.Node, len(bNode.Entries))
	copy(entries, bNode.Entries)
	// reset the entries of bNode to nil
	bNode.Entries = nil

	// make a new node
	newNode := &model.BranchNode{}

	// pick seed1 and seed2 , the 2 initial entries that each will be inserted into different nodes
	entries, seed1, seed2 := pickSeeds(entries)
	bNode.Insert(seed1)
	bNode.UpdateRectangle()

	newNode.Insert(seed2)
	newNode.UpdateRectangle()
	// then insert the rest according to insertNext until nothing left to insert
	for len(entries) > 0 {
		var pickedEntry, insertedNodeIntf model.Node
		entries, pickedEntry, insertedNodeIntf = pickNext(entries, bNode, newNode)
		insertedNode := insertedNodeIntf.(*model.BranchNode)
		insertedNode.Insert(pickedEntry)
	}
	return bNode, newNode
}

func splitLeaf(node model.Node) (model.Node, model.Node) {
	lNode := node.(*model.LeafNode)
	// save its entries into another list
	entries := make([]*model.LeafEntry, len(lNode.Entries))
	copy(entries, lNode.Entries)
	// reset the entries of lNode to nil
	lNode.Entries = nil

	// make a new node
	newNode := &model.LeafNode{}

	// pick seed1 and seed2 , the 2 initial entries that each will be inserted into different nodes
	entries, seed1, seed2 := pickSeedsLeaf(entries)
	// fmt.Println(entries, *seed1, *seed2)
	lNode.Insert(seed1)
	lNode.UpdateRectangle()

	newNode.Insert(seed2)
	newNode.UpdateRectangle()
	// then insert the rest according to insertNext until nothing left to insert
	var pickedEntry *model.LeafEntry
	var insertedNode *model.LeafNode
	for len(entries) > 0 {
		entries, pickedEntry, insertedNode = pickNextLeaf(entries, lNode, newNode)
		insertedNode.Insert(pickedEntry)
	}
	return lNode, newNode
}

func pickSeeds(entries []model.Node) ([]model.Node, model.Node, model.Node) {
	bestIdx1, bestIdx2 := 0, 0
	maxArea := math.Inf(-1)

	for idx1 := 0; idx1 < len(entries); idx1++ {
		rect1 := entries[idx1].GetRectangle()
		for idx2 := idx1 + 1; idx2 < len(entries); idx2++ {

			rect2 := entries[idx2].GetRectangle()
			newRect := rectangle.GetCombinedRectangle(rect1, rect2)
			// fmt.Println(*newRect.LowerLeft)
			// fmt.Println(*newRect.UpperRight)
			// fmt.Println(newRect.Area)
			if newRect.Area > maxArea {
				bestIdx1, bestIdx2 = idx1, idx2
				maxArea = newRect.Area
			}
		}
	}

	entry1, entry2 := entries[bestIdx1], entries[bestIdx2]
	// delete the seeds from the entries for it is already chosen
	// delete from bestIdx2 first for it is the latter, so to maintain original order
	entries = append(entries[:bestIdx2], entries[bestIdx2+1:]...)
	entries = append(entries[:bestIdx1], entries[bestIdx1+1:]...)

	return entries, entry1, entry2
}

func pickNext(entries []model.Node, node1, node2 model.Node) ([]model.Node, model.Node, model.Node) {
	var bestD1, bestD2 float64
	maxDif := math.Inf(-1)
	bestIdx := 0

	// find d1 and d2 for each remaining entry
	// d1 and d2 are the increase of area if entry E is inserted to node1 and node2
	// and note the entry E which results the maximum difference of d1 and d2
	for idx := range entries {
		rect := entries[idx].GetRectangle()
		rect1 := node1.GetRectangle()
		rect2 := node2.GetRectangle()

		newRect1 := rectangle.GetCombinedRectangle(rect, rect1)
		d1 := math.Abs(newRect1.Area - rect1.Area)

		newRect2 := rectangle.GetCombinedRectangle(rect, rect2)
		d2 := math.Abs(newRect2.Area - rect2.Area)

		dif := math.Abs(d1 - d2)
		if dif > maxDif {
			bestIdx, bestD1, bestD2, maxDif = idx, d1, d2, dif
		}
	}

	pickedEntry := entries[bestIdx]

	// pick the node to insert the best entry
	var insertedNode model.Node
	if bestD1 < bestD2 {
		insertedNode = node1
	} else {
		insertedNode = node2
	}
	// delete the said entry
	entries = append(entries[:bestIdx], entries[bestIdx+1:]...)

	return entries, pickedEntry, insertedNode
}

func pickSeedsLeaf(entries []*model.LeafEntry) ([]*model.LeafEntry, *model.LeafEntry, *model.LeafEntry) {

	inpEntries := make([]model.Node, len(entries))
	for idx := range entries {
		inpEntries[idx] = entries[idx]
	}

	entriesIntf, seedIntf1, seedIntf2 := pickSeeds(inpEntries)
	var retEntries []*model.LeafEntry
	for idx := range entriesIntf {
		retEntries = append(retEntries, entriesIntf[idx].(*model.LeafEntry))
	}
	seed1, seed2 := seedIntf1.(*model.LeafEntry), seedIntf2.(*model.LeafEntry)
	return retEntries, seed1, seed2
}

func pickNextLeaf(entries []*model.LeafEntry, node1, node2 model.Node) ([]*model.LeafEntry, *model.LeafEntry, *model.LeafNode) {

	inpEntries := make([]model.Node, len(entries))
	for idx := range entries {
		inpEntries[idx] = entries[idx]
	}

	entriesIntf, pickedEntryIntf, insertedNodeIntf := pickNext(inpEntries, node1, node2)
	var retEntries []*model.LeafEntry
	for idx := range entriesIntf {
		retEntries = append(retEntries, entriesIntf[idx].(*model.LeafEntry))
	}

	pickedEntry, insertedNode := pickedEntryIntf.(*model.LeafEntry), insertedNodeIntf.(*model.LeafNode)
	return retEntries, pickedEntry, insertedNode
}
