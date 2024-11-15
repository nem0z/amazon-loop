package topK

import (
	"container/list"
	"sync"
	"time"

	containers "github.com/nem0z/amazon-loop/heap"
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
	pq       containers.PriorityQueue
	rpq      containers.ReversePriorityQueue
	freqs    map[int]int
	products *list.List
	mu       sync.Mutex
}

func New(k int) *TopKService {
	return &TopKService{
		k:        k,
		pq:       containers.NewPriorityQueue(),
		rpq:      containers.NewReversePriorityQueue(),
		freqs:    map[int]int{},
		products: list.New(),
		mu:       sync.Mutex{},
	}
}

func (topK *TopKService) Freqs() map[int]int {
	return topK.freqs
}

func (topK *TopKService) Push(product Product) {
	topK.mu.Lock()
	defer topK.mu.Unlock()

	product.timestamp = time.Now()

	topK.products.PushFront(product)
	topK.freqs[product.id]++

	pqItem := &containers.PriorityQueueItem{
		Id:   product.id,
		Freq: topK.freqs[product.id],
	}

	if topK.pq.Len() < topK.k {
		topK.pq.Push(pqItem)
	} else {
		if topK.pq.Update(pqItem.Id, pqItem.Freq) {
			return
		}

		topK.rpq.Push(pqItem)
		for topK.pq.Peek().(*containers.PriorityQueueItem).Freq < topK.rpq.Peek().(*containers.PriorityQueueItem).Freq {
			rpqMax := topK.rpq.Pop()
			pqMin := topK.pq.Pop()

			topK.pq.Push(rpqMax)
			topK.rpq.Push(pqMin)
		}
	}
}

func (topK *TopKService) Update() {

	for {
		back := topK.products.Back()
		if back == nil {
			return
		}

		product := back.Value.(Product)
		if time.Since(product.timestamp) < time.Second*3 {
			break
		}

		topK.products.Remove(back)
		topK.freqs[product.id]--

		topK.pq.Update(product.id, topK.freqs[product.id])
		topK.rpq.Update(product.id, topK.freqs[product.id])
	}

	for topK.pq.Peek().(*containers.PriorityQueueItem).Freq < topK.rpq.Peek().(*containers.PriorityQueueItem).Freq {
		rpqMax := topK.rpq.Pop()
		pqMin := topK.pq.Pop()

		topK.pq.Push(rpqMax)
		topK.rpq.Push(pqMin)
	}
}

func (topK *TopKService) Lock() {
	topK.mu.Lock()
}

func (topK *TopKService) Unlock() {
	topK.mu.Unlock()
}

func (topK *TopKService) Collect() ([]*containers.PriorityQueueItem, []*containers.PriorityQueueItem) {
	topK.Update()
	return topK.pq.Collect(), topK.rpq.Collect()
}
