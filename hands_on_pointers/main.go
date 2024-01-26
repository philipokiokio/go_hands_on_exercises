package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"sort"
)

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and i am walking")
}
func (d dog) run() {
	d.first = "Alexis"
	fmt.Println("My name is", d.first, "and i am running")
}

type youngin interface {
	walk()
	run()
}

func YoungRun(y youngin) {
	y.run()
}

type person struct {
	First string
}

func renamer(p person, newName string) person {
	p.First = newName

	return p
}
func renamerP(p *person, name string) {
	p.First = name
}

func main() {
	x := 42

	fmt.Println("This is the memory address for x", &x)

	y := &x
	fmt.Println("This is the defrenced pointer of y which is x: ", *y, "memory location", y)

	//	Hands on Exercise 76
	dog1 := dog{"Roger"}
	dog1.walk()
	dog1.run()
	dog2 := dog{"spike"}
	dog2.walk()
	dog2.run()

	YoungRun(dog2)
	YoungRun(dog1)
	//	####################
	philip := person{First: "Philip"}
	new_philip := renamer(philip, "reggie")
	fmt.Println("Value did not change here is philip's name", philip.First, "and new_philip's name", new_philip.First)
	woman := &person{"xxx"}
	renamerP(woman, "Lexus")
	fmt.Println("woman's name is", woman.First)

	//##################### Hands On JSON ##########################
	var persons []person

	data, err := json.Marshal([]person{philip, *woman, new_philip})

	if err != nil {
		log.Fatalln("Error Marshalling struct", err)

	}
	os.Stdout.Write(data)
	fmt.Println()

	//[{"First":"Philip"},{"First":"Lexus"},{"First":"reggie"}]
	var personBlob []byte = []byte(`[{"First":"Philip"},{"First":"Lexus"},{"First":"reggie"}]`)

	unmarshalErr := json.Unmarshal(personBlob, &persons)

	if err != nil {
		fmt.Println("error", unmarshalErr)

	}
	fmt.Printf("%+v\n", persons)

	//	####################### sort ############################

	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Miss Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("___________________")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println("_________ Hands on Bcrypt __________")

	rawPassword := "irregularGuy"
	bs, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("hashed password: ", string(bs), "raw password:", rawPassword)

	err = bcrypt.CompareHashAndPassword(bs, []byte(rawPassword))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Logged In")

}
