package rtree

import (
	"errors"

	"github.com/r-tree/pkg/coordinate"
	"github.com/r-tree/pkg/rectangle"
)

// Remove remove entry on index from the tree
func (rtree *RTree) Remove(index int, coord *coordinate.Coordinate) error {

	rect := rectangle.NewRectangle(coord, coord)
	containingLeaf := rtree.FindLeaf(rtree.Root, index, rect)

	removedIdx := -1
	for idx := range containingLeaf.Entries {
		entry := containingLeaf.Entries[idx]
		if entry.Index == index {
			removedIdx = idx
			break
		}
	}

	if removedIdx == -1 {
		return errors.New("Entry not found")
	}

	containingLeaf.Entries = append(containingLeaf.Entries[:removedIdx], containingLeaf.Entries[removedIdx+1:]...)
	rtree.CondenseTree(containingLeaf)
	return nil

}
