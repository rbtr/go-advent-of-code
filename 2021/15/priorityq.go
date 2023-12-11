package main

import (
	"container/heap"
)

// An Vert is something we manage in a priority queue.
type Vert struct {
	entrycost int
	totalcost int
	coord     XY
	index     int
}

// A priorityq implements heap.Interface and holds Items.
type priorityq []*Vert

func (pq priorityq) Len() int { return len(pq) }

func (pq priorityq) Less(i, j int) bool {
	return pq[i].totalcost < pq[j].totalcost
}

func (pq priorityq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityq) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Vert)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityq) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *priorityq) update(item *Vert) {
	heap.Fix(pq, item.index)
}
