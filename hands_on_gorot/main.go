package main

import "fmt"

type person struct {
	FirstName string
}

func (p *person) speak() {
	fmt.Println("Hi, i am", p.FirstName)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	philip := person{FirstName: "Philip"}
	philip.speak()
	//saySomething(philip)
	//
	philip2 := &person{FirstName: "Ruger"}
	philip2.speak()

	saySomething(philip2)
}
