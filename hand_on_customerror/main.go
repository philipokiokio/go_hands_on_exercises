package main

import (
	"fmt"
	"log"
)

type customError struct {
	foo string
}

func (c *customError) Error() string {
	return fmt.Sprintln(fmt.Errorf("Custom Error Struct: %v", c.foo))
}

func main() {

	c := customError{foo: "This is dumb"}

	log.Fatalln(c.Error())

}
