package main

import (
	"fmt"
	"sync"
	"time"
)

const buffersize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == buffersize {
		b.cond.Wait()
	}
	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal()
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait()
	}

	takeout := b.items[0]
	b.items = b.items[1:]

	// fmt.Println("Consumed:", takeout)
	b.cond.Signal()
	return takeout
}

func producer(b *buffer, wg *sync.WaitGroup){
	defer wg.Done()

	for i := range 10 {
		b.produce(i + 100)
		time.Sleep(time.Millisecond * 100)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup){
	defer wg.Done()

	for range 10 {
		consumed := b.consume()
		fmt.Println("Consumed:", consumed)
		time.Sleep(time.Millisecond * 200)
	}
}


func main() {
	buffer := newBuffer(buffersize)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
