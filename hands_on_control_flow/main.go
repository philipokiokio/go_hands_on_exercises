package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := rand.Intn(10), rand.Intn(10)

	if x < 4 && y < 4 {
		fmt.Println("X and Y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("X and Y are greater than 6")

	} else if x >= 4 && x <= 6 {
		fmt.Println("X  is greater than  or equal to 4 while less than or equal to 6")

	} else if y != 5 {
		fmt.Println("Y not 5")

	} else {
		fmt.Println("None of the conditions were met")
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		fmt.Printf("random conditional for i of value %v\n", i)
		x, y := rand.Intn(10), rand.Intn(10)

		if x < 4 && y < 4 {
			fmt.Println("X and Y are less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("X and Y are greater than 6")

		} else if x >= 4 && x <= 6 {
			fmt.Println("X  is greater than  or equal to 4 while less than or equal to 6")

		} else if y != 5 {
			fmt.Println("Y not 5")

		} else {
			fmt.Println("None of the conditions were met")
		}

	}

	for i := 0; i < 42; i++ {
		fmt.Printf("Iteration i: %v\n", i)
		switch x := rand.Intn(5); x {
		case 1:
			fmt.Println("X is of value 1")
		case 2:
			fmt.Println("X is of value 2")
		case 3:
			fmt.Println("X is of value 3")
		case 4:
			fmt.Println("X is value 4")
		default:
			fmt.Println("X is nilth value. X is 0")
		}
	}

	tye := 84
	for tye > 42 {
		fmt.Println(tye)
		tye--
	}
	fmt.Printf("Tye is now the value of happiness of value: %v\n", tye)

	for {
		if tye > 84 {
			break
		}
		fmt.Println(tye)
		tye++
	}

	for x := 0; x <= 1000; x++ {
		if x%2 != 0 {
			continue
		}
		fmt.Printf("x of value: %v is an even number\n", x)

	}

	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("loop through xi with index %d, and value %d \n", i, v)

	}

	m := map[string]int{"reggie": 42, "yaweh": 42}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m["james"])

	if value, ok := m["james"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("No value")
	}
}
