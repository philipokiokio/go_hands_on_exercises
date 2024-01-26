package main

import "fmt"

func main() {

	arr := [5]int{}

	fmt.Println("empty array:", arr)
	for i := 0; i < 5; i++ {
		arr[i] = i

	}

	for index, value := range arr {
		fmt.Printf("array index of %v \t value of %v \n", index, value)
	}
}
