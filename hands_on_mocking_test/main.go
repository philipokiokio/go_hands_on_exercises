package main

import "fmt"

func Add(x, y int) int {
	return x + y
}
func Subtract(x, y int) int {
	return x - y
}

func DoMaths(x int, y int, f func(int, int) int) int {
	return f(x, y)
}
func main() {
	sum := Add(5, 7)

	fmt.Println(sum)

	fmt.Println(DoMaths(10, 9, Subtract))
	//Hands on Annoymous Function

	anno := func(x, y int) int {
		return x * y
	}

	fmt.Println(anno(5, 4))
}
