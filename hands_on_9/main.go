package main

import "fmt"

var bobGirl string = "Can we pretend that airplanes in the night sky are shooting stars!"

const OldRoger = "John Doe"

func main() {

	fmt.Println("Package level variable", bobGirl)

	fmt.Println("Package level const", OldRoger)

	//	local scope or block scope
	bread := "This is only possible with a function"
	fmt.Println("This is a local scoped variable", bread)

}


