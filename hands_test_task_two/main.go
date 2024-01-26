package main

import (
	"fmt"
	"word_test/quot"
	"word_test/word"
)

func main() {
	fmt.Println(word.Count(quot.SunAlso))

	for k, v := range word.UseCount(quot.SunAlso) {
		fmt.Println(v, k)
	}
}
