package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where intialization of my code occurs")
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("X  is of value %v \n", x)

	if x <= 100 && x >= 0 {
		fmt.Println("X  is between 0 & 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("X  is between 101 & 200")

	} else {
		fmt.Println("X  is between 201 & 250")

	}

	switch {
	case x <= 100 && x >= 0:
		fmt.Println("X  is between 0 & 100")

	case x >= 101 && x <= 200:
		fmt.Println("X  is between 101 & 200")

	case x >= 201 && x <= 250:
		fmt.Println("X  is between 201 & 250")

	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
