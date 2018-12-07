package model

type RTree struct {
	MinEntry, MaxEntry int
	Root               *BranchNode
}
