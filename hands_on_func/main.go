package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("my name is %s and i am %d years old\n", p.first, p.age)
}

type shape interface {
	area() float64
}
type square struct {
	length int
	width  int
}

func (s square) area() float64 {
	return float64(s.length * s.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func Info(s shape) {
	fmt.Println(s.area())
}

func main() {
	happiness := foo()

	notHappiness, data := bar()

	fmt.Println(happiness)
	fmt.Println(notHappiness, data)

	//	Hands on 59

	fmt.Println("############### 59 ###################")
	ranger := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//Varidic sum
	fmt.Println(varidicSum(ranger...))
	//slice sum
	fmt.Println(sum(ranger))
	fmt.Println("______________________________________")

	//	HANDS ON 60

	fmt.Println("############### 60 ###################")
	defer fmt.Println("This was deferred")
	fmt.Println("This printline should work before the defer")
	fmt.Println("Hands on 60 divider was deffered")

	fmt.Println("______________________________________")
	//	HANDS ON 61
	fmt.Println("############### 61 ###################")

	philip := person{first: "Philip", age: 26}

	philip.speak()
	fmt.Println("______________________________________")

	//	HANDS ON 62
	fmt.Println("############### 62 ###################")

	sq := square{length: 54, width: 20}
	ci := circle{radius: 64}

	Info(sq)
	Info(ci)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "randall"
}

// ############### 59 ###################
func varidicSum(num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

func sum(num []int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

// ______________________________________
