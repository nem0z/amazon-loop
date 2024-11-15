package containers

import (
	"container/heap"
)

type PriorityQueueItem struct {
	Id    int
	Freq  int
	Index int
}

type PriorityQueue struct {
	slice []*PriorityQueueItem
	m     map[int]*PriorityQueueItem
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{
		slice: []*PriorityQueueItem{},
		m:     map[int]*PriorityQueueItem{},
	}
}

func (pq *PriorityQueue) Len() int {
	return len(pq.slice)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.slice[i].Freq < pq.slice[j].Freq
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.slice[i], pq.slice[j] = pq.slice[j], pq.slice[i]
	pq.slice[i].Index = i
	pq.slice[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	newItem := x.(*PriorityQueueItem)

	if item, ok := pq.m[newItem.Id]; ok {
		item.Freq = newItem.Freq
		heap.Fix(pq, item.Index)
	} else {
		newItem.Index = pq.Len()
		pq.slice = append(pq.slice, newItem)
		pq.m[newItem.Id] = newItem
		heap.Fix(pq, newItem.Index)
	}
}

func (pq *PriorityQueue) Pop() any {
	n := len(pq.slice) - 1
	pq.Swap(0, n) // Swap root with last element
	item := pq.slice[n]
	pq.slice = pq.slice[:n] // Remove last element
	delete(pq.m, item.Id)
	heap.Fix(pq, 0) // Heapify to maintain heap property
	return item
}

func (pq *PriorityQueue) Peek() any {
	return pq.slice[0]
}

func (pq *PriorityQueue) Update(id int, freq int) bool {
	if item, ok := pq.m[id]; ok {
		item.Freq = freq
		heap.Fix(pq, item.Index)
		return true
	}

	return false
}

func (pq *PriorityQueue) Collect() []*PriorityQueueItem {
	// items := make([]int, pq.Len())
	// for i := range pq.slice {
	// 	items[i] = pq.slice[i].Id
	// }
	// return items

	return pq.slice
}

func (pq *PriorityQueue) MinFreq() int {
	return pq.slice[0].Id
}
