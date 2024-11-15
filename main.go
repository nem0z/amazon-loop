package main

import (
	"fmt"
	"time"

	topKService "github.com/nem0z/amazon-loop/topK"
)

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

			topK.Lock()
			fmt.Println("---------------------")
			min, max := topK.Collect()
			fmt.Println(topK.Freqs())

			for i := range min {
				fmt.Printf("%v => %v (%v) - ", min[i].Id, min[i].Freq, min[i].Index)
			}
			fmt.Println()

			for i := range max {
				fmt.Printf("%v => %v (%v) - ", max[i].Id, max[i].Freq, max[i].Index)
			}
			fmt.Println()
			fmt.Println("---------------------")
			topK.Unlock()
		}
	}()

	select {}
}
