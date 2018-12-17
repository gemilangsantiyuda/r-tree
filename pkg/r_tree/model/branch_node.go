package model

import (
	"errors"

	"github.com/r-tree/pkg/rectangle"
)

// BranchNode struct for branch nodes
type BranchNode struct {
	Parent    *BranchNode
	Rectangle *rectangle.Rectangle
	Entries   []Node
}

// GetRectangle function so it implements Node
func (bn *BranchNode) GetRectangle() *rectangle.Rectangle {
	return bn.Rectangle
}

// UpdateRectangle updating rectangle as combined area of its entries' rectangles
func (bn *BranchNode) UpdateRectangle() {
	// if it has no entries, just return
	if len(bn.Entries) == 0 {
		return
	}

	// else create new rectangle out of combined rectangles from all its entries
	newRect := &rectangle.Rectangle{}
	*newRect = *bn.Entries[0].GetRectangle()

	for idx := range bn.Entries {
		entry := bn.Entries[idx]
		rect := entry.GetRectangle()
		newRect = rectangle.GetCombinedRectangle(newRect, rect)
	}

	bn.Rectangle = newRect
}

// Insert insert nodes to the branch node's entries
func (bn *BranchNode) Insert(node Node) {
	node.SetParent(bn)
	bn.Entries = append(bn.Entries, node)
	bn.UpdateRectangle()
}

// GetParent get this node's parent
func (bn *BranchNode) GetParent() Node {
	return bn.Parent
}

// SetParent set this node's parent
func (bn *BranchNode) SetParent(node Node) {
	bn.Parent = node.(*BranchNode)
}

// RemoveChild remove child
func (bn *BranchNode) RemoveChild(child Node) error {

	removedIdx := -1
	for idx := range bn.Entries {
		if bn.Entries[idx] == child {
			removedIdx = idx
			break
		}
	}

	if removedIdx == -1 {
		err := errors.New("Child not found")
		return err
	}

	bn.Entries = append(bn.Entries[:removedIdx], bn.Entries[removedIdx+1:]...)
	return nil
}
