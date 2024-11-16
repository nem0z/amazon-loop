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

			ids, freqs := topK.Collect()
			fmt.Println(ids)
			}
		}
	}()

	select {}
}
