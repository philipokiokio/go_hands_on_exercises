package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	//bs, err := json.Marshal(p1)
	//if err != nil {
	//	fmt.Println(err)
	//	//Log fatal is another option that can be used.
	//	return
	//}
	//fmt.Println(string(bs))

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)

	}
	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Struct to JSON failed in toJSON func. Here is error:%v", err)

	}
	return bs, err
}
