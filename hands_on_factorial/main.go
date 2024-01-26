package main

import (
	"fmt"
	"math/big"
)

func main() {
	value := int64(100)
	recieverChan := factorial(value)

	factorialData := factorTotalizer(recieverChan)

	fmt.Println("Factorial Function completed")
	fmt.Printf("Factorial Function for %d is %s\n", value, factorialData.String())
	fmt.Println("Program has ended")
}

func factorial(n int64) <-chan int64 {
	biChan := make(chan int64)

	go func() {
		for i := n; i > 0; i-- {
			biChan <- i
		}
		close(biChan)
	}()
	return biChan
}
func factorTotalizer(reChan <-chan int64) *big.Int {
	totalFactorial := big.NewInt(1)

	for value := range reChan {
		v := big.NewInt(value)

		totalFactorial.Mul(v, totalFactorial)
	}
	return totalFactorial
}
