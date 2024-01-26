package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42

	}()

	fmt.Println(<-c)

	bc := make(chan int, 1)
	bc <- 21

	fmt.Println(<-bc)
	c1 := make(chan int)

	go gen(c1)
	receive(c1)

	fmt.Println("about to exit")
}

func gen(c1 chan int) {

	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func receive(c <-chan int) {
	for value := range c {
		fmt.Println("channel produces:", value)

	}
}
