package main

import "fmt"

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

func main() {

	person1 := person{first: "Vigilante", last: "Random", favoriteIceCream: []string{"vanila", "chocolate", "strawberry"}}
	person2 := person{first: "rango", last: "Jezzos", favoriteIceCream: []string{"mint", "chocolate", "vanila-choco"}}

	fmt.Println(person1.first, person1.last)
	for _, value := range person1.favoriteIceCream {
		fmt.Printf("Person1 ice cream: %v\n", value)
	}
	fmt.Println("_______________________")

	fmt.Println(person2.first, person2.last)
	for _, value := range person2.favoriteIceCream {
		fmt.Printf("Person2 ice cream: %v\n", value)
	}

	//	Hands on 54
	fmt.Println("########### Hands on 54 ###############")

	peopleDB := make(map[string]person)

	peopleDB[person1.last] = person1
	peopleDB[person2.last] = person2

	for key, value := range peopleDB {
		for i := 0; i < len(value.favoriteIceCream); i++ {
			fmt.Printf("%s likes %s \n", key, value.favoriteIceCream[i])
		}
	}

	fmt.Println("_______________________")

	//	Hands on 55
	fmt.Println("########### Hands on 55 ###############")

	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine engine
		make   string
		model  string
		doors  int
		color  string
	}

	electricEngine := engine{electric: true}
	combustionEngine := engine{electric: false}

	tesla := vehicle{engine: electricEngine, make: "Tesla", model: "One", doors: 4, color: "red"}
	hyundia := vehicle{engine: combustionEngine, make: "Hyundia", model: "Sonata", doors: 4, color: "metallic"}

	fmt.Println(tesla.engine.electric, tesla.make, tesla.model, tesla.doors, tesla.color)
	fmt.Println(hyundia.engine.electric, hyundia.make, hyundia.model, hyundia.doors, hyundia.color)

	fmt.Println("_______________________")

	//	Hands on 56
	fmt.Println("########### Hands on 56 ###############")

	philip := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Philip", friends: map[string]int{
			"God": 1,
		},
		favDrinks: []string{"Yorgurt"},
	}
	fmt.Println(philip.favDrinks, philip.first, philip.friends)
}
