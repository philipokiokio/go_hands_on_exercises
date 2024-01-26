package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type User struct {
	First string
	Age   int
}

type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByFirst []User

func (a ByFirst) Len() int           { return len(a) }
func (a ByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirst) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	//jsonUser, err := json.Marshal(users)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(jsonUser))
	err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {

		fmt.Println(err)
		return
	}

	userString := []byte(`[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`)

	var userRef []User
	err = json.Unmarshal(userString, &userRef)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", userRef)
	fmt.Println("________________ SORT _______________________________")
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println("________________ SORT BY _______________________________")
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByFirst(users))
	fmt.Println(users)

}
