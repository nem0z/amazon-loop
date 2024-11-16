package topK

import (
	"time"

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
