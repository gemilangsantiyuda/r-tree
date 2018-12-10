package model

// RTree struct to store the r-tree and its constraints
type RTree struct {
	MinEntry, MaxEntry int
	Root               Node
}
