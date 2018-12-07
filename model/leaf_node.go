package model

import (
	repo "github.com/r-tree/repository"
)
type LeafNode struct {
	Rectangle Rectangle
	Entries   []*LeafEntry
}

func (ln *LeafNode) UpdateRectangle() {

	NewRectangle := ln.Entries[0].Rectangle
	for _, entry := range ln.Entries {
		NewRectangle = repo.
	}

}
