package pkgapi

import (
	"container/heap"
	"testing"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() (v any) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func TestHeap(t *testing.T) {
	var h = &IntHeap{1, 2, 4, 0, 5}
	heap.Init(h)

	heap.Push(h, 3)

	for h.Len() > 0 {
		t.Logf("%d \n", heap.Pop(h))
	}
}

type Item struct {
	v        string
	i        int
	priority int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Swap(i, j int) {
	queue := *pq
	queue[i], queue[j] = queue[j], queue[i]

	queue[i].i = i
	queue[i].i = j
}

func (pq *PriorityQueue) Less(i, j int) bool {
	queue := *pq
	return queue[i].priority > queue[j].priority
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.i = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak

	item.i = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
