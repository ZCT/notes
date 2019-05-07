package main

import (
	"fmt"
	"sync"
)

func main() {
	test := make(chan struct{}, 0)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			select {
			case <-test:
				fmt.Println("received closing signal")
				wg.Done()
				break
			default:
				fmt.Println("doing nothing")
			}
		}
	}()
	close(test)
	wg.Wait()
}
