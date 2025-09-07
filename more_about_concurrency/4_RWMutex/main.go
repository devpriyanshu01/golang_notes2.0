package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu    sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read Value -", counter)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int){
	defer wg.Done()
	rwmu.Lock()
	counter = value
	fmt.Println("Written Value -", value)
	rwmu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go readCounter(&wg)
	}
	time.Sleep(time.Microsecond * 50)
	wg.Add(1)
	go writeCounter(&wg, 25)	
	
	wg.Wait()
}
