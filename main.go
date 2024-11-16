package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	topKService "github.com/nem0z/amazon-loop/topK"
)

func verify(top []int, freqs map[int]int) bool {
	sortedFreq := make([]int, len(freqs))
	i := 0
	for _, freq := range freqs {
		sortedFreq[i] = freq
		i++
	}

	sort.Ints(sortedFreq)

	for _, id := range top {
		if freqs[id] < sortedFreq[len(sortedFreq)-len(top)] {
			return false
		}
	}

	return true
}

func main() {

	topK := topKService.New(3)

	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			topK.Push(topKService.GenerateProduct())
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)

			ids, freqs := topK.Collect()
			fmt.Println(ids)
			if !verify(ids, freqs) {
				log.Fatal("Invalid result :", freqs)
			}
		}
	}()

	select {}
}
