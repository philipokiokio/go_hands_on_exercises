package main

import "fmt"

func main() {
	happiness := 42

	var si = make([]int, 10)

	for i := 0; i < 10; i++ {

		si[i] = happiness + i
	}

	for index, value := range si {
		fmt.Printf("Slice index: %v, \t Value: %v \t Type: %T\n", index, value, value)
	}
	var x []int
	x = si
	fmt.Println(x)
	//	Hands on 44 Challenge
	fmt.Println("_______HANDS ON 44_________")

	fmt.Println(si[:5])
	fmt.Println(si[5:])
	fmt.Println(si[2:7])
	fmt.Println(si[1:6])

	//	Hands on 45 Challenge
	fmt.Println("_______HANDS ON 45_________")
	si = append(si, 52)
	fmt.Println(si)
	si = append(si, 53, 54, 55)
	fmt.Println(si)
	y := []int{56, 57, 58, 59, 60}
	si = append(si, y...)
	fmt.Println(si)

	//	Hands on 46
	fmt.Println("_______HANDS ON 45_________")
	x = append(x[:3], x[6:]...)

	fmt.Println(x)

	//	Hands on 47

	states := []string{
		"Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois", "Indiana",
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"Montana", "Nebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania", "Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington",
		"West Virginia",
		"Wisconsin",
		"Wyoming",
	}
	fmt.Printf("States has a len of %v and a cap of %v\n", len(states), cap(states))
	//style 1
	i := 0
	for {
		if i >= len(states) {
			break
		}

		fmt.Printf("State index %v,\t with value: %v\n", i, states[i])
		i++
	}
	//	style 2
	for i := 0; i < len(states); i++ {

		fmt.Printf("State index %v,\t with value: %v\n", i, states[i])
	}

	//	Hands on 48
	fmt.Println("_______HANDS ON 48_________")

	mdsi := make([][]string, 2, 2)

	mdsi[0] = []string{"James", "Bond", "Shaken not stirred"}
	mdsi[1] = []string{"miss", "Moneypenny", "008"}

	for index, value := range mdsi {
		fmt.Printf("dimension %v\n", index)
		for internal_index, internal_value := range value {
			fmt.Printf("Outer index: %v and internal index:%v has value: %v\n", index, internal_index, internal_value)
		}
	}

}
