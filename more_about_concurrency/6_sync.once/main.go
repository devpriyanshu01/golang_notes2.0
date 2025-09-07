package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("This function will be called only once, no matter how many times it is called goroutines")
}
func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("Goroutine -", i)
			once.Do(initialize)
			// initialize()
		}()
	}
	wg.Wait()
}
