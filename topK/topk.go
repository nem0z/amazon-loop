package topK

import (
	"container/list"
	"sync"
	"time"

	"github.com/nem0z/amazon-loop/heap"
)

type topKService struct {
	k        int
	minHeap  heap.MinHeap
	maxHeap  heap.MaxHeap
	freqs    map[int]int
	products *list.List
	mu       sync.Mutex
}

func New(k int) *topKService {
	return &topKService{
		k:        k,
		minHeap:  heap.NewMinHeap(),
		maxHeap:  heap.NewMaxHeap(),
		freqs:    map[int]int{},
		products: list.New(),
		mu:       sync.Mutex{},
	}
}

func (topK *topKService) Collect() ([]int, map[int]int) {
	topK.mu.Lock()
	defer topK.mu.Unlock()
	topK.Update(true)

	topKitems := topK.minHeap.Collect()
	ids := make([]int, len(topKitems))

	for i := range topKitems {
		ids[i] = topKitems[i].Id
	}

	return ids, topK.freqs
}

func (topK *topKService) Freqs() map[int]int {
	return topK.freqs
}

func (topK *topKService) balance() {
	for topK.minHeap.Peek().(*heap.HeapItem).Freq < topK.maxHeap.Peek().(*heap.HeapItem).Freq {
		topK.minHeap.Push(topK.maxHeap.Pop())
		topK.maxHeap.Push(topK.minHeap.Pop())
	}
}

func (topK *topKService) Push(product Product) {
	topK.mu.Lock()
	defer topK.mu.Unlock()

	product.timestamp = time.Now()

	topK.products.PushFront(product)
	topK.freqs[product.id]++

	heapItem := &heap.HeapItem{
		Id:   product.id,
		Freq: topK.freqs[product.id],
	}

	if topK.minHeap.Len() < topK.k {
		topK.minHeap.Push(heapItem)
		return
	}

	if topK.minHeap.Update(heapItem.Id, heapItem.Freq) {
		return
	}

	topK.maxHeap.Push(heapItem)
	topK.balance()
}

func (topK *topKService) Update(collecting bool) {
	if !collecting {
		topK.mu.Lock()
		defer topK.mu.Unlock()
	}

	for back := topK.products.Back(); back != nil; back = topK.products.Back() {
		product := back.Value.(Product)
		if time.Since(product.timestamp) < time.Second*3 {
			break
		}

		topK.products.Remove(back)
		topK.freqs[product.id]--

		topK.minHeap.Update(product.id, topK.freqs[product.id])
		topK.maxHeap.Update(product.id, topK.freqs[product.id])
	}

	topK.balance()
}
