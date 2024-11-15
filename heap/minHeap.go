package heap

import stdHeap "container/heap"

type MinHeap struct {
	slice []*HeapItem
	m     map[int]*HeapItem
}

func NewMinHeap() MinHeap {
	return MinHeap{
		slice: []*HeapItem{},
		m:     map[int]*HeapItem{},
	}
}

func (heap *MinHeap) Len() int {
	return len(heap.slice)
}

func (heap *MinHeap) Less(i, j int) bool {
	return heap.slice[i].Freq < heap.slice[j].Freq
}

func (heap *MinHeap) Swap(i, j int) {
	heap.slice[i], heap.slice[j] = heap.slice[j], heap.slice[i]
	heap.slice[i].Index = i
	heap.slice[j].Index = j
}

func (heap *MinHeap) Push(x any) {
	newItem := x.(*HeapItem)

	if item, ok := heap.m[newItem.Id]; ok {
		item.Freq = newItem.Freq
		stdHeap.Fix(heap, item.Index)
	} else {
		newItem.Index = heap.Len()
		heap.slice = append(heap.slice, newItem)
		heap.m[newItem.Id] = newItem
		stdHeap.Fix(heap, newItem.Index)
	}
}

func (heap *MinHeap) Pop() any {
	n := len(heap.slice) - 1
	heap.Swap(0, n)
	item := heap.slice[n]
	heap.slice = heap.slice[:n]
	delete(heap.m, item.Id)
	stdHeap.Fix(heap, 0)
	return item
}

func (heap *MinHeap) Peek() any {
	return heap.slice[0]
}

func (heap *MinHeap) Update(id int, freq int) bool {
	if item, ok := heap.m[id]; ok {
		item.Freq = freq
		stdHeap.Fix(heap, item.Index)
		return true
	}

	return false
}

func (heap *MinHeap) Collect() []*HeapItem {
	// items := make([]int, heap.Len())
	// for i := range heap.slice {
	// 	items[i] = heap.slice[i].Id
	// }
	// return items

	return heap.slice
}

func (heap *MinHeap) MinFreq() int {
	return heap.slice[0].Id
}
