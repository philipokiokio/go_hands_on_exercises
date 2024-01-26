package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(2)
	go conclevel()

	go rocklevel()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func conclevel() {
	for i := 0; i < 100; i++ {
		fmt.Println("step", i)
	}
	wg.Done()
}

func rocklevel() {

	fmt.Println("go")
	wg.Done()
}
