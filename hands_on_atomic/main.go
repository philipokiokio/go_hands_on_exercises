package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var value int64
	value = 1
	const gs = 100
	fmt.Println("Value: ", value)

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&value, 1)
			atomic.LoadInt64(&value)
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}

	wg.Wait()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	fmt.Println("Incremented value: ", value)
}
