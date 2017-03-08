package main

import (
	"fmt"
	"sync"
)

type mutexCounter struct {
	mu sync.Mutex
	x  int64
}

func (c *mutexCounter) Add(x int64) {
	c.mu.Lock()
	c.x += x
	c.mu.Unlock()
}

func (c *mutexCounter) Value() (x int64) {
	c.mu.Lock()
	x = c.x
	c.mu.Unlock()
	return
}

func main() {
	counter := mutexCounter{}
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				counter.Add(1)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(counter.Value())
}
