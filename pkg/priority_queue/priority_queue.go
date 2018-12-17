package priorityQueue

import "container/heap"

// PriorityQueue struct that wraps priorityqueue <- which is the one that implements golang.heap
type PriorityQueue struct {
	pq priorityqueue
}

// New PriorityQueue (wrapper of real priority queue)
func New() PriorityQueue {
	pq := make(priorityqueue, 0)
	heap.Init(&pq)
	PQ := PriorityQueue{
		pq: pq,
	}
	return PQ
}

// Push an object into the PriorityQueue
func (PQ PriorityQueue) Push(node Node, priority float64) {
	item := &Item{
		Object:   node,
		Priority: priority,
	}
	heap.Push(&PQ.pq, item)
}

// Pop the nearest node
func (PQ PriorityQueue) Pop() (Node, float64) {
	item := heap.Pop(PQ.&pq).(*Item)
	return (item.Object, item.Priority)
}

type priorityqueue []*Item

func (pq priorityqueue) Len() int {
	return len(pq)
}

func (pq priorityqueue) Less(i, j int) bool {
	// We want Pop to give us the  lowest priority (nearest).
	return pq[i].priority < pq[j].priority
}

func (pq priorityqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityqueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
