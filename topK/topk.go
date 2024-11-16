package topK

import (
	"container/list"
	"sync"
	"time"

	"github.com/nem0z/amazon-loop/heap"
	"golang.org/x/exp/rand"
)

type Product struct {
	id        int
	timestamp time.Time
}

func GenerateProduct() Product {
	return Product{
		id: rand.Intn(10),
	}
}

type TopKService struct {
	k        int
	minHeap  heap.MinHeap
	maxHeap  heap.MaxHeap
	freqs    map[int]int
	products *list.List
	mu       sync.Mutex
}

func New(k int) *TopKService {
	return &TopKService{
		k:        k,
		minHeap:  heap.NewMinHeap(),
		maxHeap:  heap.NewMaxHeap(),
		freqs:    map[int]int{},
		products: list.New(),
		mu:       sync.Mutex{},
	}
}

func (topK *TopKService) Freqs() map[int]int {
	return topK.freqs
}

func (topK *topKService) balance() {
	for topK.minHeap.Peek().(*heap.HeapItem).Freq < topK.maxHeap.Peek().(*heap.HeapItem).Freq {
		topK.minHeap.Push(topK.maxHeap.Pop())
		topK.maxHeap.Push(topK.minHeap.Pop())
	}
}
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
